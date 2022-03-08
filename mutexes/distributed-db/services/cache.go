package services

import (
	"log"
	"strings"
	"sync"
	"time"

	"distributed-db/models"
)

type CacheRepository interface {
	Get(keys []int) []models.CacheItem
	Set(items map[int]models.CacheItem)
	Delete(keys []int)
	GetAllKeys() []int
}

type HTTPClient interface {
	Get(node string, keys []string) ([]models.CacheItem, error)
	Set(node, key, value string) (models.CacheItem, error)
	SetBatch(node string, items map[int]models.CacheItem) ([]models.CacheItem, error)
	Gossip(node string, newNodes models.NodesMap, tokensChecksum string) (oldNodes models.NodesMap, err error)
	Tokens(node string) (models.TokenMappings, error)
}

func NewCache(cacheRepo CacheRepository, httpClient HTTPClient, tokens *models.Tokens) CacheSvc {
	return CacheSvc{
		cacheRepo:  cacheRepo,
		httpClient: httpClient,
		tokens:     tokens,
		//hashCache => local cache for generated hashes and the server they belong to
		// save a bit of computational time
	}
}

type CacheSvc struct {
	cacheRepo  CacheRepository
	httpClient HTTPClient
	tokens     *models.Tokens
}

func (svc CacheSvc) Get(keys []string) []models.CacheItem {
	// there's a problem here when having foreign records (stolen tokens) on current node
	// lookup current node anyways
	keyToNode := map[string]string{}
	tokenToNode := map[int]string{}
	for _, key := range keys {
		token := int(models.HashKey(key))
		node := svc.tokens.GetNode(token)
		tokenToNode[token] = node
		keyToNode[key] = node
	}

	nodeToTokens := map[string][]int{}
	for token, node := range tokenToNode {
		nodeToTokens[node] = append(nodeToTokens[node], token)
	}
	nodeToKeys := map[string][]string{}
	for key, node := range keyToNode {
		nodeToKeys[node] = append(nodeToKeys[node], key)
	}

	cacheItems := make([]models.CacheItem, 0)
	for node, tokens := range nodeToTokens {
		if node == svc.tokens.Nodes.Current() {
			items := svc.cacheRepo.Get(tokens)
			for _, item := range items {
				item.Node = node
				cacheItems = append(cacheItems, item)
			}
			continue
		}

		nodeKeys := nodeToKeys[node]
		items, err := svc.httpClient.Get(node, nodeKeys)
		if err != nil {
			log.Printf("could not get cache items from node: %s, %v", node, err)
			continue
		}

		for _, item := range items {
			item.Node = node
			cacheItems = append(cacheItems, item)
		}
	}

	// TOKEN STEALING
	// SET and SET BATCH
	// add foreign tokens to gossip data
	// once the data is successfully streamed remove it from foreign tokens

	// GET
	// when getting data from a coordinator:
	// 1. check local node
	// 2. check the node which has foreign token
	// 3. check the node which contains the data
	// also when making request verify the node status

	return cacheItems
}

func (svc CacheSvc) Set(key, value string) (models.CacheItem, error) {
	token := int(models.HashKey(key))
	node := svc.tokens.GetNode(token)
	if node == svc.tokens.Nodes.Current() {
		item := models.CacheItem{
			Key:       key,
			Value:     value,
			UpdatedAt: time.Now().UTC(),
		}

		svc.cacheRepo.Set(map[int]models.CacheItem{token: item})
		item.Node = node

		return item, nil
	}

	// save the record no matter what and retry in the background
	// if retry has failed x amount of times, apply token range stealing
	// if retry has succeeded, remove the item from the current node
	return svc.httpClient.Set(node, key, value)
}

func (svc CacheSvc) SetBatch(items map[int]models.CacheItem) []models.CacheItem {
	resItems := make([]models.CacheItem, 0)
	localItems := map[int]models.CacheItem{}
	nodesToForeignItems := map[string]map[int]models.CacheItem{}
	setNode := func(items []models.CacheItem, node string) {
		for _, item := range items {
			item.Node = node
		}
	}

	// split items into local items and foreign items
	// local items => belong to current node
	// map of nodes to foreign items => belong to different nodes
	for token, item := range items {
		node := svc.tokens.GetNode(token)
		if node == svc.tokens.Nodes.Current() {
			localItems[token] = item
			item.Node = node
			resItems = append(resItems, item)
			continue
		}

		nodesToForeignItems[node][token] = item
	}
	// save local items on the current node
	if len(localItems) > 0 {
		svc.cacheRepo.Set(localItems)
	}

	// attempt to save foreign items on each node they belong to
	for node, foreignItems := range nodesToForeignItems {
		batchItems, err := svc.httpClient.SetBatch(node, foreignItems)
		setNode(batchItems, node)
		if err != nil {
			// if batch call failed, save the items on the current node
			// they will get redistributed by the streamer worker anyways
			log.Printf("could not set batch for node %s: %v", node, err)
			svc.cacheRepo.Set(foreignItems)
			svc.tokens.SetForeignTokens(foreignItems, svc.tokens.Nodes.Current())
			setNode(batchItems, svc.tokens.Nodes.Current())
		}

		resItems = append(resItems, batchItems...)
	}

	return resItems
}

func (svc CacheSvc) Delete(keys []string) {
	// also implement retry mechanism
	// implement tombstone and make it short-lived
	// think about if necessary to gossip tombstone
}

func (svc CacheSvc) Gossip() {
	nodes := svc.tokens.Nodes.ListActive(2)
	if len(nodes) == 0 {
		return
	}

	log.Println("gossiping to:", strings.Join(nodes, ","))
	for _, node := range nodes {
		oldNodes, err := svc.httpClient.Gossip(node, svc.tokens.Nodes.Map(), svc.tokens.Checksum())
		if err != nil {
			log.Printf("could not make http call for gossip: %v", err)
			svc.tokens.Nodes.Fail(node)
			continue
		}

		svc.tokens.Nodes.Set(oldNodes)
	}
}

func (svc CacheSvc) GetTokens() map[int]string {
	return svc.tokens.Mappings
}

func (svc CacheSvc) UpdateTokens(node string, newNodes models.NodesMap, tokensChecksum string) (models.NodesMap, error) {
	svc.tokens.Nodes.Gossip(node)
	svc.tokens.Nodes.Set(newNodes)

	if svc.tokens.Checksum() != tokensChecksum {
		tokens, err := svc.httpClient.Tokens(node)
		if err != nil {
			return models.NodesMap{}, err
		}
		svc.tokens.Merge(tokens)
	}

	return svc.tokens.Nodes.Map(), nil
}

func (svc CacheSvc) Stream(retryBatches map[string]map[int]models.CacheItem) map[string]map[int]models.CacheItem {
	keys := svc.cacheRepo.GetAllKeys()

	// LOOKUP NEW ITEMS
	tryingToStream, nodeToBatches := 0, retryBatches
	for _, token := range keys {
		node := svc.tokens.GetNode(token)
		if node != svc.tokens.Nodes.Current() {
			continue
		}

		items := svc.cacheRepo.Get([]int{token})
		if len(items) == 0 {
			continue
		}

		if nodeToBatches[node] == nil {
			nodeToBatches[node] = map[int]models.CacheItem{}
		}
		nodeToBatches[node][token] = items[0]
		tryingToStream++
	}
	if tryingToStream > 0 {
		log.Printf("trying to stream %d new item(s)", tryingToStream)
	}

	// PREPARE BATCHES
	type batch struct {
		node         string
		items        map[int]models.CacheItem
		keysToDelete []int
	}
	batchSize := 10
	batches := make([]batch, 0) // => maybe make it a map[string][]batch => now we can't access it without for looping
	for node, items := range nodeToBatches {
		i, batchItems, keysToDelete := 0, map[int]models.CacheItem{}, make([]int, 0)
		for token, item := range items {
			if i == batchSize {
				b := batch{
					node:         node,
					items:        batchItems,
					keysToDelete: keysToDelete,
				}
				batches = append(batches, b)
				i, batchItems, keysToDelete = 0, map[int]models.CacheItem{}, []int{}
			}
			i++
			keysToDelete = append(keysToDelete, token)
			batchItems[token] = item
		}
		if len(batchItems) > 0 {
			b := batch{
				node:         node,
				items:        batchItems,
				keysToDelete: keysToDelete,
			}
			batches = append(batches, b)
		}
	}

	// START BATCH STREAMING
	failedToStream, failedBatches := 0, map[string]map[int]models.CacheItem{}
	failBatch := func(node string, batchItems map[int]models.CacheItem) {
		for token, item := range batchItems {
			if failedBatches[node] == nil {
				failedBatches[node] = map[int]models.CacheItem{}
			}
			failedBatches[node][token] = item
			failedToStream++
		}
	}

	var wg sync.WaitGroup
	maxConcurrentBatches := 10
	for i := 0; i < len(batches); i += maxConcurrentBatches {

	}

	for i, b := range batches {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := svc.httpClient.SetBatch(b.node, b.items)
			if err != nil {
				failBatch(b.node, b.items)
			}
			svc.cacheRepo.Delete(b.keysToDelete)
			log.Printf("successfully streamed %d items", b.keysToDelete)
		}()
		if i%maxConcurrentBatches == 0 {
			wg.Wait()
		}
	}

	if failedToStream > 0 {
		log.Printf("failed to stream %d items", failedToStream)
	}
	return failedBatches
}

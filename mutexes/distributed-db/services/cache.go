package services

import (
	"fmt"
	"log"
	"strings"

	"distributed-db/models"
)

type CacheRepository interface {
	Get(keys []string) []models.CacheItem
	Set(key, value string) models.CacheItem
}

type HTTPClient interface {
	Set(node, key, value string) (models.CacheItem, error)
	Get(node string, keys []string) ([]models.CacheItem, error)
	Gossip(node string, newNodes models.NodesMap, tokensChecksum string) (oldNodes models.NodesMap, err error)
	Tokens(node string) (models.TokenMappings, error)
}

func NewCache(cacheRepo CacheRepository, httpClient HTTPClient, tokens *models.Tokens) CacheSvc {
	return CacheSvc{
		cacheRepo:  cacheRepo,
		httpClient: httpClient,
		tokens:     tokens,
	}
}

type CacheSvc struct {
	cacheRepo  CacheRepository
	httpClient HTTPClient
	tokens     *models.Tokens
}

func (svc CacheSvc) Get(keys []string) []models.CacheItem {
	keyToNode := map[string]string{}
	sumToNode := map[string]string{}
	for _, key := range keys {
		sum := fmt.Sprintf("%d", models.HashKey(key))
		node := svc.tokens.GetNode(key)
		sumToNode[sum] = node
		keyToNode[key] = node
	}

	nodeToSums := map[string][]string{}
	for sum, node := range sumToNode {
		nodeToSums[node] = append(nodeToSums[node], sum)
	}
	nodeToKeys := map[string][]string{}
	for key, node := range keyToNode {
		nodeToKeys[node] = append(nodeToKeys[node], key)
	}

	cacheItems := make([]models.CacheItem, 0)
	for node, sums := range nodeToSums {
		if node == svc.tokens.Nodes.Current() {
			items := svc.cacheRepo.Get(sums)
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
		}

		for _, item := range items {
			item.Node = node
			cacheItems = append(cacheItems, item)
		}
	}

	return cacheItems
}

func (svc CacheSvc) Set(key, value string) (models.CacheItem, error) {
	node := svc.tokens.GetNode(key)
	if node == svc.tokens.Nodes.Current() {
		item := svc.cacheRepo.Set(key, value)
		item.Node = node
		return item, nil
	}

	return svc.httpClient.Set(node, key, value)
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

func (svc CacheSvc) GetTokens() map[int]string {
	return svc.tokens.Mappings
}

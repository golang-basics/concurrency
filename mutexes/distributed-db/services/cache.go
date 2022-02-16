package services

import (
	"log"
	"strings"

	"distributed-db/models"
)

type CacheRepository interface {
	Get(key string) *models.CacheItem
	GetMany(keys []string) []models.CacheItem
	Set(key, value string)
}

type HTTPClient interface {
	Get(node string, key string) (models.CacheItem, error)
	Gossip(node string, nodes []string, tokensChecksum string) (oldNodes []string, err error)
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
	return svc.cacheRepo.GetMany(keys)
}

func (svc CacheSvc) Set(key, value string) {
	svc.cacheRepo.Set(key, value)
}

func (svc CacheSvc) Gossip() {
	nodes := svc.tokens.Nodes.List(2)
	log.Println("gossiping to:", strings.Join(nodes, ","))

	allNodes := append([]string{}, svc.tokens.Nodes.List(len(svc.tokens.Nodes.Map))...)
	allNodes = append(allNodes, svc.tokens.Nodes.CurrentNode)
	for _, node := range nodes {
		oldNodes, err := svc.httpClient.Gossip(node, allNodes, svc.tokens.Checksum())
		if err != nil {
			log.Printf("could not make http call for gossip: %v", err)
			continue
		}

		svc.tokens.Nodes.Add(oldNodes...)
	}
}

func (svc CacheSvc) UpdateTokens(node string, newNodes []string, tokensChecksum string) ([]string, error) {
	oldNodes := svc.tokens.Nodes.List(len(svc.tokens.Nodes.Map))
	if svc.tokens.Checksum() != tokensChecksum {
		svc.tokens.Nodes.Add(newNodes...)
		tokens, err := svc.httpClient.Tokens(node)
		if err != nil {
			return []string{}, err
		}
		svc.tokens.Merge(tokens)
	}
	return oldNodes, nil
}

func (svc CacheSvc) GetTokens() map[int]string {
	return svc.tokens.Mappings
}

package services

import (
	"fmt"
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
	Gossip(node string, nodes []string, tokensChecksum string) error
	Tokens(node string) (models.TokenMappings, error)
}

func NewCache(cacheRepo CacheRepository, httpClient HTTPClient, tokens models.Tokens) CacheSvc {
	return CacheSvc{
		cacheRepo:  cacheRepo,
		httpClient: httpClient,
		tokens:     tokens,
	}
}

type CacheSvc struct {
	cacheRepo  CacheRepository
	httpClient HTTPClient
	tokens     models.Tokens
}

func (svc CacheSvc) Get(keys []string) []models.CacheItem {
	return svc.cacheRepo.GetMany(keys)
}

func (svc CacheSvc) Set(key, value string) {
	svc.cacheRepo.Set(key, value)
}

func (svc CacheSvc) UpdateTokens(node string, newNodes []string, tokensChecksum string) ([]string, error) {
	oldNodes := svc.tokens.Nodes.List(len(svc.tokens.Nodes.Map))
	for _, n := range newNodes {
		svc.tokens.Nodes.Add(n)
	}

	if svc.tokens.Checksum() != tokensChecksum {
		tokens, err := svc.httpClient.Tokens(node)
		if err != nil {
			return []string{}, err
		}
		//svc.tokens.Mappings = tokens
		// try also adding the current node
		svc.tokens.AddNode(node)
		fmt.Println(tokens)
		fmt.Println(svc.tokens.Mappings)
		// for some reason this duplicated data for the new server
		//svc.tokens.AddNode(node)
	}
	return oldNodes, nil
}

func (svc CacheSvc) Gossip() {
	nodes := svc.tokens.Nodes.List(2)
	log.Println("gossiping to:", strings.Join(nodes, ","))

	allNodes := append([]string{}, svc.tokens.Nodes.List(len(svc.tokens.Nodes.Map))...)
	allNodes = append(allNodes, svc.tokens.Nodes.CurrentNode)
	for _, node := range nodes {
		if err := svc.httpClient.Gossip(node, allNodes, svc.tokens.Checksum()); err != nil {
			log.Printf("could not make http call for gossip: %v", err)
			continue
		}
	}
}

func (svc CacheSvc) GetTokens() map[uint64]string {
	return svc.tokens.Mappings
}

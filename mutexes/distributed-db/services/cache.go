package services

import (
	"crypto/md5"
	"fmt"
	"log"
	"strings"

	"distributed-db/models"
)

type CacheRepository interface {
	Get(key string) *models.CacheItem
	GetMany(keys []string) []models.CacheItem
	Set(key, value string)
	GetSummary() models.Summary
}

type HTTPClient interface {
	Gossip(peer string, summary models.Summary) error
	Get(peer string, key string) (models.CacheItem, error)
}

func NewCache(cacheRepo CacheRepository, httpClient HTTPClient, peers models.Peers) CacheSvc {
	return CacheSvc{
		cacheRepo:  cacheRepo,
		httpClient: httpClient,
		peers:      peers,
	}
}

type CacheSvc struct {
	cacheRepo  CacheRepository
	httpClient HTTPClient
	peers      models.Peers
}

func (svc CacheSvc) Get(keys []string) []models.CacheItem {
	return svc.cacheRepo.GetMany(keys)
}

func (svc CacheSvc) Set(key, value string) {
	svc.cacheRepo.Set(key, value)
}

func (svc CacheSvc) GossipSummary() {
	summary := svc.cacheRepo.GetSummary()
	if len(summary) == 0 {
		return
	}

	peers := svc.peers.List(models.MinimumPeers)
	log.Println("gossiping to:", strings.Join(peers, ","))
	for _, peer := range peers {
		if err := svc.httpClient.Gossip(peer, summary); err != nil {
			log.Printf("could not make http call for gossip: %v", err)
			continue
		}
	}
}

func (svc CacheSvc) ResolveSummary(peer string, summary models.Summary) {
	svc.peers.Add(peer)

	for key, updatedAt := range summary {
		oldCacheItem := svc.cacheRepo.Get(key)
		if oldCacheItem != nil && updatedAt.Sub(oldCacheItem.UpdatedAt) < 0 {
			continue
		}

		newCacheItem, err := svc.httpClient.Get(peer, key)
		if err != nil {
			log.Printf("could not make http call for get(%s): %v", key, err)
			continue
		}

		if oldCacheItem == nil {
			fmt.Println("brand new item, does not exist in current db")
			svc.Set(key, newCacheItem.Value)
			continue
		}

		oldSum := fmt.Sprintf("%x", md5.Sum([]byte(oldCacheItem.Value)))
		newSum := fmt.Sprintf("%x", md5.Sum([]byte(newCacheItem.Value)))
		if oldSum == newSum {
			continue
		}
		fmt.Println("old item that has updated contents")
		svc.Set(key, newCacheItem.Value)
	}
}

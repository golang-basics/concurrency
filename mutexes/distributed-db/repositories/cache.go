package repositories

import (
	"fmt"
	"sync"
	"time"

	"distributed-db/models"
)

func NewCache() *Cache {
	return &Cache{
		data:    map[string]models.CacheItem{},
	}
}

type Cache struct {
	mu      sync.RWMutex
	data    map[string]models.CacheItem
}

func (c *Cache) Set(key, value string) models.CacheItem {
	c.mu.Lock()
	defer c.mu.Unlock()

	item := models.CacheItem{
		Key:       key,
		Value:     value,
		UpdatedAt: time.Now().UTC(),
	}

	sum := fmt.Sprintf("%d", models.HashKey(key))
	c.data[sum] = item
	return item
}

func (c *Cache) Get(keys []string) []models.CacheItem {
	c.mu.RLock()
	defer c.mu.RUnlock()

	items := make([]models.CacheItem, 0)
	for _, key := range keys {
		item, ok := c.data[key]
		if ok {
			items = append(items, item)
		}
	}

	return items
}

package repositories

import (
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

func (c *Cache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	cacheValue := models.CacheItem{
		Key:       key,
		Value:     value,
		UpdatedAt: time.Now().UTC(),
	}
	c.data[key] = cacheValue
}

func (c *Cache) Get(key string) *models.CacheItem {
	c.mu.RLock()
	defer c.mu.RUnlock()

	val, ok := c.data[key]
	if !ok {
		return nil
	}

	return &val
}

func (c *Cache) GetMany(keys []string) []models.CacheItem {
	c.mu.RLock()
	defer c.mu.RUnlock()

	values := make([]models.CacheItem, 0)
	for _, k := range keys {
		v, ok := c.data[k]
		if ok {
			values = append(values, v)
		}
	}

	return values
}

package repositories

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"distributed-db/models"
)

func NewCache(dataDir string) *Cache {
	cache := &Cache{
		data:    map[string]models.CacheItem{},
		dataDir: dataDir,
	}
	cache.init()
	return cache
}

type Cache struct {
	mu      sync.RWMutex
	data    map[string]models.CacheItem
	dataDir string
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

func (c *Cache) Snapshot() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	err := os.MkdirAll(c.dataDir, os.ModePerm)
	if err != nil && os.IsNotExist(err) {
		log.Printf("could not create directory: %v", err)
		return err
	}

	file, err := os.Create(fmt.Sprintf("%s/db.json", c.dataDir))
	if err != nil {
		log.Printf("could not create file: %v", err)
		return err
	}

	err = json.NewEncoder(file).Encode(c.data)
	if err != nil {
		log.Printf("could not encode cache data to file: %v", err)
		return err
	}
	return nil
}

func (c *Cache) init() {
	dbFile, err := os.Open(fmt.Sprintf("%s/db.json", c.dataDir))
	if err != nil {
		return
	}

	var cacheData map[string]models.CacheItem
	err = json.NewDecoder(dbFile).Decode(&cacheData)
	if err != nil {
		log.Printf("could not decode database file: %v", err)
		return
	}

	c.data = cacheData
}

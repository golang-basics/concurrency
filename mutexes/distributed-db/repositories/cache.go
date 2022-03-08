package repositories

import (
	"fmt"
	"log"
	"os"
	"sync"

	"distributed-db/models"
	"encoding/json"
)

func NewCache(dataDir string) *Cache {
	cache := &Cache{
		data:    map[int]models.CacheItem{},
		dataDir: dataDir,
	}
	cache.init()
	return cache
}

type Cache struct {
	mu      sync.RWMutex
	data    map[int]models.CacheItem
	dataDir string
}

func (c *Cache) Get(keys []int) []models.CacheItem {
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

func (c *Cache) GetAllKeys() []int {
	keys := make([]int, 0, len(c.data))
	for key := range c.data {
		keys = append(keys, key)
	}
	return keys
}

// make this a batch function that accepts multiple items
// make it accept key and item instead
func (c *Cache) Set(items map[int]models.CacheItem) {
	for key, item := range items {
		c.mu.Lock()
		c.data[key] = item
		c.mu.Unlock()
	}
}

func (c *Cache) Delete(keys []int) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for _, key := range keys {
		delete(c.data, key)
	}
}

// Snapshot stores the in-memory database to db.json file
// CRITICAL: This is very INEFFICIENT and only for development purposes.
// A better approach would be only keeping a small portion of data in-memory,
// the rest to be kept in small chunks per file
// and a background compaction process that merges and optimizes data
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

// init creates the initial state for the database from db.json file
// CRITICAL: This is very INEFFICIENT and only for development purposes.
// A better approach would be to use the disk directory and avoid initializing
// the database like this. Creating optimized indexes and using file seeking can
// drastically improve performance
func (c *Cache) init() {
	dbFile, err := os.Open(fmt.Sprintf("%s/db.json", c.dataDir))
	if err != nil {
		return
	}

	var cacheData map[int]models.CacheItem
	err = json.NewDecoder(dbFile).Decode(&cacheData)
	if err != nil {
		log.Printf("could not decode database file: %v", err)
		return
	}

	c.data = cacheData
}

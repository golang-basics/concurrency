package repositories

import (
	"sync"
	"time"

	"distributed-db/models"
)

func NewCache() *Cache {
	return &Cache{
		data:    map[string]models.CacheItem{},
		summary: models.Summary{},
	}
}

type Bucket struct {
	ID int
	CreatedAt time.Time
	Data map[string]models.CacheItem
}

type Cache struct {
	mu      sync.RWMutex
	data    map[string]models.CacheItem // the current bucket, overwrite after a new bucket is created
	summary models.Summary //only used to gossip newly created items

	index map[string]int // keys to bucket ids (only populate the index once the bucket is saved in the list of buckets)
	buckets map[int]Bucket // all buckets (only append to this list, once a new bucket is created)
}

func (c *Cache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	// check against lastCreatedAt
	// if bigger than X duration, create a new bucket
	// save the old bucket in the background

	// if the key is also found inside index
	// update index and update the value inside its bucket

	// Maybe GET RID of the "data" field and use "buckets" only

	cacheValue := models.CacheItem{
		Key:       key,
		Value:     value,
		UpdatedAt: time.Now().UTC(),
	}
	c.data[key] = cacheValue

	// change the format to: [key:bucket][timestamp]
	c.summary[key] = cacheValue.UpdatedAt
}

func (c *Cache) Get(key string) *models.CacheItem {
	c.mu.RLock()
	defer c.mu.RUnlock()

	// first check the current bucket
	// if the key is present, return the item

	// if key is not present, look up the index
	// then return the item from bucket

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

func (c *Cache) GetSummary() models.Summary {
	c.mu.Lock()
	defer c.mu.Unlock()
	defer func() {
		c.summary = models.Summary{}
	}()

	return c.summary
}

type Index struct {
	bucketID int // last created at => lastBucket.CreatedAt
	lastCreatedAt map[int]time.Time // bucket:createdAt
	buckets []int // hint to know how many buckets are in total without loading them all
	data map[string]int // key:bucketID
}

func (c *Cache) Init() {
	// keep a small number of buckets in memory
}

func (c *Cache) Snapshot() {
	// save the index to a file
	// save each bucket in its own file
}

// GOSSIP ONLY spreads information about the nodes
// DATA REPLICATION is done through REPLICATION FACTOR and CONSISTENCY LEVEL (WRITE/READ CONSISTENCY LEVEL)

// For a really efficient database like Cassandra
// we need to implement COMPACTION and SSTables and have an INDEX and MemTable
// COMPACTION (runs in the background):
// Merge multiple SSTables into bigger SSTables -> update the index
// SSTables:
// SSTables are sorted string tables stored on Disk once the MemTable is flushed.
// When MemTable is flushed -> create/update the INDEX
// MemTable:
// Represents In-Memory data of the database
// INDEX:
// Represents a map[Key:SSTable Offset] (kept on disk)
// The INDEX is kept on disk because its size can be big
// SUMMARY
// Represents an index for the INDEX or a map[N Keys Range (Bucket) : INDEX FILE] => map[0:index1], map[100:index2], map[200:index3]
// COMMIT LOG
// Append only file in case the MemTable gets lost. It is destroyed after the MemTable is flushed

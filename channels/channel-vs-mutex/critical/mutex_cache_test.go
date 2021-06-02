package critical

import "sync"

type mutexCache struct {
	cache map[string]string
	mu    sync.Mutex
}

func newMutexCache() *mutexCache {
	return &mutexCache{
		cache: map[string]string{},
	}
}

func (c *mutexCache) set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = value
}

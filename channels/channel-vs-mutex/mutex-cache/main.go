package main

import (
	"fmt"
	"sync"
)

// try running the example like: go run -race main.go
func main() {
	c := newMutexCache()
	var wg sync.WaitGroup

	wg.Add(5)
	go func() {
		c.set("k", "v1")
		wg.Done()
	}()
	go func() {
		c.set("k", "v2")
		wg.Done()
	}()
	go func() {
		c.set("k", "v3")
		wg.Done()
	}()
	go func() {
		c.set("k", "v4")
		wg.Done()
	}()
	go func() {
		fmt.Println(c.get("k"))
		wg.Done()
	}()
	wg.Wait()
	c.set("kn", "vn")

	fmt.Println(c.get("k"))
	fmt.Println(c.get("kn"))
}

type mutexCache struct {
	cache map[string]string
	mu    sync.RWMutex
}

func newMutexCache() mutexCache {
	return mutexCache{
		cache: map[string]string{},
	}
}

func (c *mutexCache) set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = value
}

func (c *mutexCache) get(key string) string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.cache[key]
}

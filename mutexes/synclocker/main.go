package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	mu := &myLock{}
	// mu := &sync.Mutex{}
	// rwMu := &sync.RWMutex{}
	// calls RLock and RUnlock inside Lock and Unlock
	// mu := rwMu.RLocker()
	c := newCache(mu)

	wg.Add(3)
	go func() {
		defer wg.Done()
		c.set("k", "v1")
	}()
	go func() {
		defer wg.Done()
		c.set("k", "v2")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("value", c.get("k"))
	}()
	wg.Wait()
}

type myLock struct {
	sync.Mutex
}

func (l *myLock) Lock() {
	fmt.Println("locking")
	l.Mutex.Lock()
}

func (l *myLock) Unlock() {
	fmt.Println("unlocking")
	l.Mutex.Unlock()
}

type cache struct {
	mu    sync.Locker
	cache map[string]string
}

func newCache(l sync.Locker) cache {
	return cache{
		mu:    l,
		cache: map[string]string{},
	}
}

func (c *cache) set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = value
}

func (c *cache) get(key string) string {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.cache[key]
}

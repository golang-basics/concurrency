package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	d := newDB()
	var wg sync.WaitGroup
	now := time.Now()
	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Println("before getting k1")
		d.get("k1")
	}()
	time.Sleep(time.Second)
	go func() {
		defer wg.Done()
		fmt.Println("before setting k1")
		d.set("k1", "v1")
	}()
	wg.Wait()
	fmt.Println("elapsed", time.Since(now))
}

func newDB() db {
	return db{values: map[string]string{}}
}

type db struct {
	values map[string]string
	mu     sync.RWMutex
}

func (d *db) set(key, value string) {
	d.mu.Lock()
	fmt.Println("setting", key)
	time.Sleep(2 * time.Second)
	defer d.mu.Unlock()
	d.values[key] = value
}

func (d *db) get(key string) string {
	d.mu.RLock()
	fmt.Println("getting", key)
	time.Sleep(10 * time.Second)
	defer d.mu.RUnlock()
	return d.values[key]
}

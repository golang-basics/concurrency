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
	wg.Add(5)

	// 1 WRITE => waiting 2s
	go func() {
		defer wg.Done()
		fmt.Println("before setting k")
		d.set("k", "v")
	}()

	// simulate order => ignored in the elapsed time
	time.Sleep(time.Second)

	// 3 READS => waiting 5s, NOT 15s,
	// because they all start the Sleep() call at the same time
	go func() {
		defer wg.Done()
		fmt.Println("before getting k1")
		d.get("k1")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("before getting k2")
		d.get("k2")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("before getting k3")
		d.get("k3")
	}()

	// simulate order => ignored in the elapsed time
	time.Sleep(time.Second)

	// 1 WRITE => waiting 2s
	go func() {
		defer wg.Done()
		fmt.Println("before setting k1")
		d.set("k1", "v1")
	}()

	// we only wait 9s instead of 19s
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
	time.Sleep(5 * time.Second)
	defer d.mu.RUnlock()
	return d.values[key]
}

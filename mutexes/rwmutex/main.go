package main

import (
	"fmt"
	"sync"
)

// make sure the run the below program using the -race flag
// go run -race main.go
func main() {
	d := newDB()
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		d.set("k1", "v1")
		wg.Done()
	}()
	go func() {
		d.set("k1", "v2")
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(d.get("k1"))
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
	defer d.mu.Unlock()
	d.values[key] = value
}

func (d *db) get(key string) string {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.values[key]
}

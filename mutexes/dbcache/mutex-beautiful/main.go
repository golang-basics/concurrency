package main

import (
	"fmt"
	"sync"
)

// to check for race conditions use the -race flag
// go run -race main.go
func main() {
	d := newDB()
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		// setting k1 here
		d.set("k1", "v1")
		wg.Done()
	}()
	go func() {
		// setting k1 as well here
		d.set("k1", "v2")
		wg.Done()
	}()
	// trying to get a value while it's being set by the go routines
	fmt.Println(d.get("k1"))
	wg.Wait()
}

func newDB() db {
	return db{values: map[string]string{}}
}

type db struct {
	values map[string]string
	mu     sync.Mutex
}

func (d *db) set(key, value string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.values[key] = value
}

func (d *db) get(key string) string {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.values[key]
}

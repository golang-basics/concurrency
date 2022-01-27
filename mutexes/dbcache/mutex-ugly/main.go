package main

import (
	"fmt"
	"sync"
)

// to check for race conditions use the -race flag
// go run -race main.go
func main() {
	d := newDB()
	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		// setting k1 here
		mu.Lock()
		d.set("k1", "v1")
		mu.Unlock()
		wg.Done()
	}()
	go func() {
		// setting k1 as well here
		mu.Lock()
		d.set("k1", "v2")
		mu.Unlock()
		wg.Done()
	}()
	// trying to get a value while it's being set by the go routines
	mu.Lock()
	fmt.Println(d.get("k1"))
	mu.Unlock()
	wg.Wait()
}

func newDB() db {
	return db{values: map[string]string{}}
}

type db struct {
	values map[string]string
}

func (d *db) set(key, value string) {
	d.values[key] = value
}

func (d db) get(key string) string {
	return d.values[key]
}

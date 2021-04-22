package main

import (
	"fmt"
	"sync"
)

func main() {
	//m := sync.Map{}
	//p := sync.Pool{}
	//o := sync.Once{}
	//c := sync.Cond{}

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
	mu sync.RWMutex
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

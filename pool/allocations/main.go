package main

import (
	"fmt"
	"sync"
)

func main() {
	var objectsCreated int
	pool := &sync.Pool{
		New: func() interface{} {
			objectsCreated++
			// 1KB slice of byte
			mem := make([]byte, 1024)
			return &mem
		},
	}

	// seed the pool with 4KB
	pool.Put(pool.New())
	pool.Put(pool.New())
	pool.Put(pool.New())
	pool.Put(pool.New())

	var wg sync.WaitGroup
	for i := 0; i < 1024*1024; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mem := pool.Get().(*[]byte)
			defer pool.Put(mem)
		}()
	}

	wg.Wait()
	fmt.Println("number of created objects in pool:", objectsCreated)
}

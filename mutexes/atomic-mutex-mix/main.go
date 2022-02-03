package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// try running this with the -race flag
// go run -race main.go
func main() {
	var count int32
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 10000; i++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			atomic.AddInt32(&count, 1)
		}()
		go func() {
			defer wg.Done()
			mu.Lock()
			count++
			mu.Unlock()
		}()
	}

	wg.Wait()
	// the result is correct, but Go is not able to properly detect a race condition
	// due to mixed usage of Mutex and Atomics
	fmt.Println("count:", count)
}

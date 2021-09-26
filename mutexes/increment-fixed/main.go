package main

import (
	"fmt"
	"sync"
)

// Try running the following program with the -race flag
// go run -race main.go
// The order does not matter at all, the correctness of our concurrent code does.
// Mutexes and Locks are about Correctness of Execution, not about preserving order
func main() {
	var count int
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(20)
	for i := 0; i < 20; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			count++
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println("count", count)
}

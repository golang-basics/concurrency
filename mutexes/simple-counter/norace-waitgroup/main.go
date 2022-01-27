package main

import (
	"fmt"
	"sync"
)

// try running the program using the -race flag
// go run -race main.go
func main() {
	var count int
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		// even if this code executes concurrently,
		// it will perform worse than normal sequential code
		// because of the concurrency overhead and unnecessary wait time
		wg.Add(1)
		go func() {
			count++
			wg.Done()
		}()
		wg.Wait()
	}
	fmt.Println("count", count)
}

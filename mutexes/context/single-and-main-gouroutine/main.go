package main

import (
	"fmt"
	"sync"
)

// try running this with the -race flag
// go run -race main.go
func main() {
	var wg sync.WaitGroup
	var count int

	wg.Add(1)
	go func() {
		defer wg.Done()
		// G1 will attempt a read operation
		//fmt.Println(count)
		// G1 will attempt a write operation
		count++
	}()
	// Main G will also attempt a write operation
	count++

	wg.Wait()
	fmt.Println("count", count)
}

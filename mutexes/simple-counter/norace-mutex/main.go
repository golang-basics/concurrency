package main

import (
	"fmt"
	"sync"
)

// try running the program using the -race flag
// go run -race main.go
func main() {
	var count int
	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
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

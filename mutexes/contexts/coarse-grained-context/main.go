package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var count int

	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			var localCount int
			mu.Lock()
			localCount = count + 1
			count = localCount
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println("count", count)
}

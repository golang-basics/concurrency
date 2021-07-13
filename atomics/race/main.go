package main

import (
	"fmt"
	"sync"
)

// to test for race condition, run this example with the -race flag
// go run -race main.go
func main() {
	var count int32
	var wg sync.WaitGroup
	wg.Add(5)
	go func() {
		defer wg.Done()
		count += 10
	}()
	go func() {
		defer wg.Done()
		count -= 15
	}()
	go func() {
		defer wg.Done()
		count++
	}()
	go func() {
		defer wg.Done()
		count = 0
	}()
	go func() {
		defer wg.Done()
		count = 100
	}()
	wg.Wait()

	fmt.Println("count", count)
}

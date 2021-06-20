package main

import (
	"fmt"
	"sync"
)

// try running the program using the -race flag
// go run -race main.go
func main() {
	var wg sync.WaitGroup
	count := 0

	wg.Add(1)
	go func() {
		defer wg.Done()
		count = 1
	}()
	wg.Wait()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("count:", count)
	}()
	wg.Wait()
}

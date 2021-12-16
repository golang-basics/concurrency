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
		// Only the G1 go routine will do the write operation, thus it's safe
		count++
	}()
	// Main G does not read or write while G1 is running
	// So this will not run in a race condition

	wg.Wait()
	fmt.Println("count", count)
}

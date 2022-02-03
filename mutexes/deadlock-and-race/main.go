package main

import (
	"fmt"
	"sync"
)

// Having both deadlock and race
// will result in the program hanging
// when ran with -race flag
// go run -race main.go => will show races, but hang forever

// Try running the program first without -race => fix the deadlock
// go run main.go
// then using the -race flag => fix the race condition
// go run -race main.go
func main() {
	var count int
	var mu sync.Mutex
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		count++

	}()
	go func() {
		defer wg.Done()
		count++
	}()

	wg.Wait()
	mu.Lock()
	mu.Lock()
	fmt.Println("count:", count)
}

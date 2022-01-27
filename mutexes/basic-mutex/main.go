package main

import (
	"fmt"
	"sync"
	"time"
)

// to check for race conditions use the -race flag
// go run -race main.go
func main() {
	var count int
	var mu sync.Mutex
	go func() {
		mu.Lock()
		count = 10
		mu.Unlock()
	}()

	// remember: the main function is also running as a go routine
	// a good rule of thumb is to avoid RW conflicts in the main
	// and keep all the concurrent operations in their own go routines
	mu.Lock()
	count = 15
	mu.Unlock()

	// to avoid using WaitGroup
	time.Sleep(time.Second)
	mu.Lock()
	fmt.Println("count", count)
	mu.Unlock()
}

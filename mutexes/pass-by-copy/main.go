package main

import (
	"fmt"
	"sync"
)

// The below example will result in a deadlock,
// because both calls to worker() pass a copy of Mutex.
// Thus, the Unlock() call will happen on the copy, not original mu variable
// Because of that, we have 2 consecutive calls to Lock(),
// which results in a deadlock
func main() {
	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		mu.Lock()
		worker(mu)
	}()
	go func() {
		defer wg.Done()
		mu.Lock()
		worker(mu)
	}()
	wg.Wait()
}

func worker(mu sync.Mutex) {
	fmt.Println("manipulating shared data")
	mu.Unlock()
}

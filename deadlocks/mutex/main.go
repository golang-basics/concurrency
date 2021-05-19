package main

import (
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(2)

	// this is pretty much what happens below
	// Lock gets called twice, because the lock was not released
	// mu.Lock()
	// mu.Lock()

	go func() {
		mu.Lock()
		wg.Done()
	}()
	go func() {
		time.Sleep(500 * time.Millisecond)
		mu.Lock()
		defer mu.Unlock()
		wg.Done()
	}()

	wg.Wait()
}

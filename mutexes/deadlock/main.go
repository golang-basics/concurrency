package main

import (
	"sync"
)

// Every Lock() call must have a corresponding Unlock() call
// If the method Lock() more times in a row without the lock being released
// that will result in a deadlock
// The same is true for RWMutex, any RLock or Lock MUST be followed by an Unlock() call.
// If either the read or write lock is called more times in a row
// it will also result in a deadlock
func main() {
	var mu sync.Mutex
	// mu.Lock()
	// mu.Lock()
	// var mu sync.Mutex
	// mu.RLock()
	// mu.Lock()

	adultsOnly(10, &mu)
	adultsOnly(18, &mu)
}

func adultsOnly(age int, mu *sync.Mutex) {
	mu.Lock()
	// defer mu.Unlock()
	if age < 18 {
		// this is like a small jump out the window :D
		return
	}
	mu.Unlock()
}

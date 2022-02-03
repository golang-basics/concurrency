package main

import "sync"

// A call to Unlock() must only happen after a call to Lock() has been made
// If done otherwise, it will result in a fatal error
// Any of the below lock types (mutexes) will result in a fatal error
func main() {
	var mu sync.Mutex
	mu.Unlock()

	// var rwMu sync.RWMutex
	// rwMu.Unlock()
	// rwMu.RUnlock()

	// var mu sync.RWMutex
	// l := mu.RLocker()
	// l.Unlock()
}

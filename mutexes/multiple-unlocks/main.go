package main

import "sync"

// A call to Unlock() must only happen after a call to Lock() has beeen made
// If done otherwise, it will result in a panic
func main() {
	var mu sync.Mutex
	// Lock() was not called
	mu.Unlock()
}

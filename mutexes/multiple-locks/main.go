package main

import "sync"

// Only multiple consecutive calls to Lock() on Write (exclusive) mutex
// will cause the program do deadlock, every other kind of mutex is considered
// to be a Read (shared) mutex, which will work perfectly fine
func main() {
	// deadlock -> fatal error
	var mu sync.Mutex
	mu.Lock()
	mu.Lock()

	// var rwMu sync.RWMutex
	// deadlock -> fatal error
	// rwMu.Lock()
	// rwMu.Lock()
	// no deadlock -> works fine (shared lock)
	// rwMu.RLock()
	// rwMu.RLock()

	// works fine -> equivalent to RLock (shared lock)
	// var mu sync.RWMutex
	// l := mu.RLocker()
	// l.Lock()
	// l.Lock()
}

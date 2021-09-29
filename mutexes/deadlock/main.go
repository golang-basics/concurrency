package main

import (
	"sync"
)

// Every Lock() call must have a corresponding Unlock() call
// If the method Lock() 2 times in a row without the lock being released
// that will result in a deadlock
func main() {
	var mu sync.Mutex
	//mu.Lock()
	//mu.Lock()
	adultsOnly(10, &mu)
	adultsOnly(18, &mu)
}

func adultsOnly(age int, mu *sync.Mutex) {
	mu.Lock()
	if age < 18 {
		return
	}
	mu.Unlock()
}

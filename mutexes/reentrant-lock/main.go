package main

import (
	"fmt"
	"sync"
)

// Reentrant/Recursive Lock
func main() {
	var mutex sync.Mutex
	var rwMutex sync.RWMutex
	recursive(rwMutex.RLocker(), 10)
	// recursiveRWMutex(&rwMutex, 10)
	recursive(&mutex, 10)
	// recursiveMutex(&mutex, 10)
}

// NO DEADLOCK => READ LOCK is SHARED aka can be acquired by as many go routines at a time
func recursiveRWMutex(mu *sync.RWMutex, n int) {
	if n < 1 {
		return
	}
	mu.RLock()
	defer mu.RUnlock()
	fmt.Println(n)
	recursiveRWMutex(mu, n-1)
}

// DEADLOCK => WRITE LOCK is EXCLUSIVE aka can only be acquired by a single go routine at a time
func recursiveMutex(mu *sync.Mutex, n int) {
	if n < 1 {
		return
	}
	mu.Lock()
	defer mu.Unlock()
	fmt.Println(n)
	recursiveMutex(mu, n-1)
}

func recursive(mu sync.Locker, n int) {
	if n < 1 {
		return
	}
	mu.Lock()
	defer mu.Unlock()
	fmt.Println(n)
	recursive(mu, n-1)
}

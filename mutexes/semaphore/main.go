package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var lock int32

// try running this with the -race flag
// go run -race main.go
func main() {
	var count int
	var wg sync.WaitGroup

	// if we increase the number of go routines
	// this could quickly be detected as a race condition
	// due to too many go routines' status being awake
	// Also there's no way we can control the go routines queue
	// or have access to the runtime internals
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			acquire()
			count++
			release()
		}()
	}

	wg.Wait()
	fmt.Println("count", count)
}

func acquire() {
	for atomic.CompareAndSwapInt32(&lock, 0, 1) {
	}
}

func release() {
	for atomic.CompareAndSwapInt32(&lock, 1, 0) {
	}
}

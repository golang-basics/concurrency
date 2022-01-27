package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// only 1 process can hold the resource at a time (has exclusive rights) -> non-shareable
func main() {
	mu := &sync.Mutex{}
	res := resource{mu: mu}

	// suppose there's process p_main in the operating system acquiring
	// an exclusive resource, and it does not plan on releasing it
	res.acquire("p_main")
	//res.release()

	// suppose the operating system has other 100 processes trying to
	// acquire the same resource through the lock that has not been released by p_main
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			fmt.Printf("resource acquired by process: p_%d\n", i+1)
			mu.Unlock()
		}(i)
	}
	wg.Wait()
}

type resource struct {
	mu        *sync.Mutex
	acquired  int32
	processID string
}

func (r *resource) acquire(processID string) {
	if atomic.CompareAndSwapInt32(&r.acquired, 0, 1) {
		r.mu.Lock()
		r.processID = processID
		fmt.Println("resource acquired by process:", processID)
	}
}

func (r *resource) release() {
	if atomic.CompareAndSwapInt32(&r.acquired, 1, 0) {
		r.mu.Unlock()
		fmt.Println("resource released by process:", r.processID)
	}
}

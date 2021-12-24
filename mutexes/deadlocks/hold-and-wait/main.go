package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// let's assume the condition is: in order to compute total we need the values from v1 and v2
func main() {
	var mu sync.Mutex
	var v1, v2, total int
	go func() {
		mu.Lock()
		v1 = 2
		v2 = 3
		// we could also simply not call the Unlock method
		runtime.Goexit()
		mu.Unlock()
	}()

	// wait for go routine to call Lock first
	// this is to avoid using wait groups
	time.Sleep(time.Second)
	mu.Lock()
	total = v1 + v2
	mu.Unlock()
	fmt.Println(total)
}

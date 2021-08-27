package main

import (
	"fmt"
	"sync"
)

func main() {
	cond := sync.NewCond(&sync.Mutex{})
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		cond.L.Lock()
		cond.Wait()
		fmt.Println("handler 1")
		cond.L.Unlock()
	}()
	go func() {
		defer wg.Done()
		cond.L.Lock()
		cond.Wait()
		fmt.Println("handler 2")
		cond.L.Unlock()
	}()
	wg.Wait()
	// Looks like the call to Broadcast is way too fast
	// Whoa, slow down. Looks like not all go routines are ready
	// You missed some Done() calls ;)
	cond.Broadcast()
}

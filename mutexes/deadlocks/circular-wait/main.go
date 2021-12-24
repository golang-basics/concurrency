package main

import (
	"fmt"
	"sync"
)

func main() {
	m1, m2 := &sync.Mutex{}, &sync.Mutex{}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		m1.Lock()
		fmt.Println("m1 locked")

		// m2.Unlock() never gets called => deadlock, while waiting to acquire lock
		m2.Lock()
		fmt.Println("m2 locked")
		m2.Unlock()
		fmt.Println("m2 unlocked")
		fmt.Println("done in go routine")
	}()

	m2.Lock()
	fmt.Println("m2 locked")

	// to simulate a circular wait we need to create the condition
	// that the call to Done() is not being made, thus the call to Wait will be stuck
	// causing the code below to deadlock, while waiting to acquire lock
	wg.Wait()
	// m1.Unlock() never gets called => deadlock, while waiting to acquire lock
	m1.Lock()
	fmt.Println("m1 locked")
	m1.Unlock()
	fmt.Println("m1 unlocked")
	fmt.Println("done in main")
}

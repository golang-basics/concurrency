package main

import (
	"fmt"
	"sync"
)

func main() {
	// 1. ADD number of waits
	// 2. Call Done() inside each go routine
	// 3. Call Wait() where you want to wait for the execution of all go routines

	// RULES
	// 1. Done MUST be called as many times as the number inside the Add call
	// 2. Waiting or not calling Done enough times will result in a deadlock
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		fmt.Println("1")
		wg.Done()
	}()
	go func() {
		fmt.Println("2")
		wg.Done()
	}()
	go func() {
		fmt.Println("3")
		wg.Done()
	}()
	wg.Wait()
}

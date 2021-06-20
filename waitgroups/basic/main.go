package main

import (
	"fmt"
	"sync"
)

func main() {
	// 1. Call Add() with the number of required operations to be waited on
	// 2. Call Done() inside each go routine (to be waited on)
	// 3. Call Wait() where you want to wait for the execution of all go routines

	// RULES
	// 1. Done() MUST be called as many times as Add()
	// 2. If calls to Done() are less than calls to Add(), it will result in a deadlock
	// 3. If calls to Done() are more than calls to Add, it will result in panic
	// 4. Calling Wait() without calling Add() will return immediately
	// 5. WaitGroup MUST always be passed by reference (as pointer)
	// 6. Calling another Wait() before the previous one returns results in panic
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

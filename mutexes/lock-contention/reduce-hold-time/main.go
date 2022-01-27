package main

import (
	"sync"
)

var mu sync.Mutex
var c int

func task1() {}
func task2() {}
func task3() {}
func inc()   { c++ }

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		work1()
	}()
	go func() {
		defer wg.Done()
		work2()
	}()
	wg.Wait()
}

// bad practice
func work1() {
	mu.Lock()
	defer mu.Unlock()
	task1()
	inc()
	task2() // I/O
	task3()
}

// good practice
func work2() {
	task1()
	func() {
		mu.Lock()
		defer mu.Unlock()
		inc()
	}()
	task2() // I/O
	task3()
}

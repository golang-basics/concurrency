package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		now := time.Now()
		// work: ~6s | acquire after: ~6s
		// work1(&mu)
		// work: ~6s | acquire after: ~1s
		work2(&mu)
		fmt.Println("work is done after:", time.Since(now))
	}()

	// simulate order of go routines
	time.Sleep(100 * time.Nanosecond)

	// other go routine that needs the same mutex
	go func() {
		defer wg.Done()
		now := time.Now()
		mu.Lock()
		fmt.Println("acquired lock after:", time.Since(now))
		mu.Unlock()
	}()

	wg.Wait()
}

// bad practice
func work1(mu *sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()
	task1()
	task2()
	task3()
}

// good practice
func work2(mu *sync.Mutex) {
	// let's say only task 1 works with the CRITICAL SECTION
	// every other task is just part of the work
	func() {
		mu.Lock()
		defer mu.Unlock()
		task1()
	}()
	task2()
	task3()
}

func task1() {
	time.Sleep(1 * time.Second)
}

func task2() {
	time.Sleep(2 * time.Second)
}

func task3() {
	time.Sleep(3 * time.Second)
}

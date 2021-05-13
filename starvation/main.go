package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var lock sync.Mutex

	wg.Add(2)
	go greedy(&wg, &lock)
	go polite(&wg, &lock)

	wg.Wait()
}

func greedy(wg *sync.WaitGroup, lock *sync.Mutex) {
	defer wg.Done()
	var count int
	for begin := time.Now(); time.Since(begin) < time.Second; {
		lock.Lock()
		time.Sleep(3*time.Nanosecond)
		lock.Unlock()
		count++
	}
	fmt.Println("greedy worker executed", count, "times")
}

func polite(wg *sync.WaitGroup, lock *sync.Mutex) {
	defer wg.Done()
	var count int
	for begin := time.Now(); time.Since(begin) < time.Second; {
		lock.Lock()
		time.Sleep(1*time.Nanosecond)
		lock.Unlock()
		lock.Lock()
		time.Sleep(1*time.Nanosecond)
		lock.Unlock()
		lock.Lock()
		time.Sleep(1*time.Nanosecond)
		lock.Unlock()
		count++
	}
	fmt.Println("polite worker executed", count, "times")
}

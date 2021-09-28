package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(2)
	go greedy(&wg, &mu)
	go polite(&wg, &mu)
	wg.Wait()
}

func greedy(wg *sync.WaitGroup, lock *sync.Mutex) {
	defer wg.Done()
	var count int
	for begin := time.Now(); time.Since(begin) < time.Second; {
		lock.Lock()
		time.Sleep(3 * time.Nanosecond)
		lock.Unlock()
		count++
		// will allow the processor to move on processing other go routines
		// will also avoid starvation
		//runtime.Gosched()
	}
	fmt.Println("greedy worker executed", count, "times")
}

func polite(wg *sync.WaitGroup, lock *sync.Mutex) {
	defer wg.Done()
	var count int
	for begin := time.Now(); time.Since(begin) < time.Second; {
		lock.Lock()
		time.Sleep(time.Nanosecond)
		lock.Unlock()
		lock.Lock()
		time.Sleep(time.Nanosecond)
		lock.Unlock()
		lock.Lock()
		time.Sleep(time.Nanosecond)
		lock.Unlock()
		count++
	}
	fmt.Println("polite worker executed", count, "times")
}

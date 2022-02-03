package main

import (
	"fmt"
	"sync"
	"time"
)

// Starvation or Lock Unfairness happens when 2 or more processes are trying to acquire a lock,
// and one of the processes abusively uses the lock way more than the rest of the processes,
// thus resulting in other processes to wait and be starved. This happens in Go when
// a go routine acquires the lock very frequently, while other go routines are constantly
// being parked and woken up in an attempt to acquire the lock, resulting in the frequent process to
// acquire the lock faster and winning the battle more often than it should.
// Starvation mostly happens because one of the processes, in our case go routine 2 releases the
// lock faster
// Note: go routine 1 will acquire the lock way more often in Go <= 1.8
func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(2)
	go func() {
		defer wg.Done()
		var count int
		for begin := time.Now(); time.Since(begin) < time.Second; {
			mu.Lock()
			count++
			time.Sleep(100 * time.Microsecond)
			mu.Unlock()
		}
		fmt.Println("g1 acquired:", count)
	}()
	go func() {
		defer wg.Done()
		var count int
		for begin := time.Now(); time.Since(begin) < time.Second; {
			time.Sleep(100 * time.Microsecond)
			mu.Lock()
			count++
			mu.Unlock()
		}
		fmt.Println("g2 acquired:", count)
	}()

	wg.Wait()
}

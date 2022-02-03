package main

import (
	"fmt"
	"sync"
)

// Using an RLocker for Write scenarios,
// has the same effect as not using any locks at all => race condition.
// RLocker is only useful for Read only scenarios,
// specifically if a function depends on a sync.Locker,
// allowing any kind of mutex to be passed as an argument
func main() {
	var count int
	var wg sync.WaitGroup
	var mu sync.RWMutex
	l := mu.RLocker()

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			l.Lock()
			// ++ operator does both read and write operations,
			// thus, this type of lock is not enough to prevent
			// any race conditions
			count++
			l.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println("count:", count)
}

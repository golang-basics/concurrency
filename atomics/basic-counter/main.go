package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// to test for race condition, run this example with the -race flag
// go run -race main.go
func main() {
	var count int64 = 0
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(10 * time.Millisecond)
		fmt.Println("count in go routine", atomic.LoadInt64(&count))
	}()

	wg.Add(50)
	for i := 0; i < 50; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(10 * time.Millisecond)
			atomic.AddInt64(&count, 1)
		}()
	}
	wg.Wait()
	fmt.Println("count in main", count)
}

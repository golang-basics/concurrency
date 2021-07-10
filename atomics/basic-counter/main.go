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
		time.Sleep(10 * time.Millisecond)
		fmt.Println(atomic.LoadInt64(&count))
		wg.Done()
	}()

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			time.Sleep(10 * time.Millisecond)
			atomic.AddInt64(&count, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(count)
}

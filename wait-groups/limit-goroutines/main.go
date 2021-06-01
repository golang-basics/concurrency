package main

import (
	"fmt"
	"sync"
	"time"
)

type request func()

func main() {
	requests := map[int]request{}
	for i := 1; i <= 100; i++ {
		f := func(n int) request {
			return func() {
				time.Sleep(500 * time.Millisecond)
				fmt.Println("request", n)
			}
		}
		requests[i] = f(i)
	}

	var wg sync.WaitGroup
	max, processed := 10, 0
	for _, r := range requests {
		// adjust this to process 10 requests at a time
		if processed > max-1 {
			break
		}
		wg.Add(1)
		go func(r request) {
			defer wg.Done()
			r()
		}(r)
		processed++
	}
	wg.Wait()
}

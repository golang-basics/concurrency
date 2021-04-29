package main

import (
	"sync"
	"time"
)
// go routine 1 will acquire the lock way more often in Go <=1.8
func main() {
	done := make(chan bool, 1)
	var mu sync.Mutex

	// goroutine 1
	go func() {
		var i int
		for {
			select {
			case <-done:
				return
			default:
				mu.Lock()
				i++
				time.Sleep(100 * time.Microsecond)
				mu.Unlock()
			}
		}
	}()

	// goroutine 2
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Microsecond)
		mu.Lock()
		mu.Unlock()
	}
	done <- true
	time.Sleep(time.Second)
}

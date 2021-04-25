package main

import (
	"sync"
	"time"
)

// To enable tracing on this program make sure to run the below command
// go build main.go
// GOMAXPROCS=2 GODEBUG=schedtrace=1000 ./main
func main() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			time.Sleep(time.Second)
			a := 0
			for i := 0; i < 1e6; i++ {
				a += 1
			}
			wg.Done()
		}()
	}
	wg.Wait()

	// wait for Global Run Queue
	time.Sleep(3 * time.Second)
}

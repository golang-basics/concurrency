package main

import (
	"sync"
	"time"
)

// To enable tracing on this program make sure to run the below command
// GOMAXPROCS=2 GODEBUG=schedtrace=1000,scheddetail=1 go run main.go
func main() {
	var wg sync.WaitGroup

	for i := 0;i < 2000 ;i++ {
		wg.Add(1)
		go func() {
			time.Sleep(20*time.Millisecond)
			a := 0

			for i := 0; i < 1e6; i++ {
				a += 1
			}

			wg.Done()
		}()
	}

	wg.Wait()
}

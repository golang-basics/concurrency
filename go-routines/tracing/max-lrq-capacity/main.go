package main

import (
	"sync"
	"time"
)

// To enable tracing on this program make sure to run the below commands
// go build main.go
// GOMAXPROCS=1 GODEBUG=schedtrace=1000 ./main
// GOMAXPROCS=2 GODEBUG=schedtrace=1000 ./main
// GOMAXPROCS=2 GODEBUG=schedtrace=1000,scheddetail=1 ./main
func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go work(&wg)
	}
	wg.Wait()
	// Wait to see the global run queue complete
	time.Sleep(3 * time.Second)
}

func work(wg *sync.WaitGroup) {
	time.Sleep(time.Second)
	var counter int
	for i := 0; i < 1e10; i++ {
		counter++
	}
	wg.Done()
}

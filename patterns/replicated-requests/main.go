package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	done := make(chan struct{})
	result := make(chan int)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go request(done, result, i+1, &wg)
	}

	first := <-result
	close(done)
	wg.Wait()

	fmt.Printf("received an answer from request: #%d\n", first)
}

func request(done chan struct{}, result chan<- int, id int, wg *sync.WaitGroup) {
	defer wg.Done()

	started := time.Now()
	latency := time.Duration(1+rand.Intn(5)) * time.Second
	select {
	case <-done:
		// we normally return here,
		// but the below time print is negligible
	case <-time.After(latency):
	}

	select {
	case <-done:
	// we normally return here,
	// but the below time print is negligible
	case result <- id:
	}

	took := time.Since(started)
	if took < latency {
		took = latency
	}
	fmt.Printf("request #%d took: \t%v\n", id, took)
}

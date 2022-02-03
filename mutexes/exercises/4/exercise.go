package main

import (
	"fmt"
	"sync"
	"time"
)

// EXERCISE:
// Find what's wrong in the exercise() function
// Make sure all the tests are passing
// DO NOT remove any Sleep() calls

// try running this with the -race flag
// go run -race exercise.go

// run the tests using:
// GOFLAGS="-count=1" go test .

func main() {
	now := time.Now()
	A := exercise()
	fmt.Println("A:", A)
	fmt.Println("elapsed:", time.Since(now))
}

func exercise() int {
	var A, B int
	var muA, muB sync.Mutex

	for i := 0; i < 1000; i++ {
		go func() {
			muA.Lock()
			defer muA.Unlock()
			muB.Lock()
			defer muB.Unlock()

			B++
			// simulate load
			time.Sleep(time.Millisecond)
			A++
		}()

	}

	var wg sync.WaitGroup
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			muA.Lock()
			defer muA.Unlock()
			A += 5
		}()
	}

	wg.Wait()
	return A
}

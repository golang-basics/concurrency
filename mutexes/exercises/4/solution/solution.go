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

// SOLUTION
// The problem with this code that uses 2 mutexes is the fact
// That the second mutex depends on the use of the first mutex,
// thus creating a perfect environment for lock contention.
// Since operations on A and B are independent, solving this is
// easy, all we have to do is release the mutex on A ASAP,
// thus allowing the other go routine to use it immediately
// and as a result the whole program executes faster
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
			A++
			muA.Unlock()

			muB.Lock()
			defer muB.Unlock()
			B++
			// simulate load
			time.Sleep(time.Millisecond)
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

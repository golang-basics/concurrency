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
// The problem with this implementation is the fine-grained
// defined context by the mutex, resulting in an undesired result.
// All we really need to fix the problem is limit the context
// of the mutex, which will also give us the expected result
func main() {
	now := time.Now()
	count := exercise()
	fmt.Println("count:", count)
	fmt.Println("elapsed:", time.Since(now))
}

func exercise() int {
	var A, B, count int
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			var a, b int
			mu.Lock()
			a = A
			b = B
			count += a + b
			A++
			B++
			mu.Unlock()
		}()
	}
	wg.Wait()
	return count
}

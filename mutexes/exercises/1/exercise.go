package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// EXERCISE:
// Find what's wrong in the following code
// Make sure all the tests are passing
// Make sure there's no deadlock or race condition
// Hint: Limit the go routines (100), then find and fix the deadlock first

// try running this with the -race flag
// go run -race exercise.go

// run the tests using:
// GOFLAGS="-count=1" go test -race .
func main() {
	clicks := exercise()
	fmt.Println("total clicks:", clicks.total)
}

func exercise() clickCounter {
	var mu sync.Mutex
	var wg sync.WaitGroup
	clicks := clickCounter{
		mu: &mu,
		elements: map[string]int{
			"btn1": 0,
			"btn2": 0,
		},
	}

	for i := 0; i < 1000; i++ {
		wg.Add(3)
		go func() {
			defer wg.Done()
			if len(clicks.elements) > 10 {
				fmt.Println("we got more than 10 elements")
			}
		}()
		go func() {
			defer wg.Done()
			clicks.add("btn1")
			clicks.add("btn2")
		}()
		go func() {
			defer wg.Done()
			total := atomic.LoadInt64(&clicks.total)
			if total > 20 {
				fmt.Println("something is wrong")
			}
		}()
	}

	wg.Wait()
	return clicks
}

type clickCounter struct {
	mu       *sync.Mutex
	elements map[string]int
	total    int64
}

func (c *clickCounter) add(element string) {
	c.mu.Lock()
	if c.elements[element]+1 > 10 {
		return
	}
	c.elements[element]++
	c.total++
	c.mu.Unlock()
}

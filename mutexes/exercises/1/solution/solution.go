package main

import (
	"fmt"
	"sync"
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

// SOLUTION
// The problem with this code is it has both:
// a deadlock and a race condition.
// The deadlock happens because we accidentally
// forget to call Unlock().
// The race condition happens because
// we mix and match Atomics with Mutexes
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
			clicks.mu.Lock()
			defer clicks.mu.Unlock()
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
			clicks.mu.Lock()
			defer clicks.mu.Unlock()
			if clicks.total > 20 {
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
	defer c.mu.Unlock()
	if c.elements[element]+1 > 10 {
		return
	}
	c.elements[element]++
	c.total++
}

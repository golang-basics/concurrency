package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

// EXERCISE:
// Find what's wrong in the following code
// Make sure all the tests are passing
// DO NOT remove any Sleep() calls

// try running this with the -race flag
// go run -race exercise.go

// run the tests using:
// GOFLAGS="-count=1" go test .

// SOLUTION
// The problem with this code is mainly because,
// we are using a regular Mutex, thus slowing down
// all the read operations. Using a RWMutex instead
// will give the same results at an improved speed
func main() {
	c := &catalog{data: map[string]string{
		"p1": "apples",
		"p2": "oranges",
		"p3": "grapes",
		"p4": "pineapple",
		"p5": "bananas",
	}}

	now := time.Now()
	exercise(c, "p1", "p2", "p3", "p4", "p5")
	fmt.Println("elapsed:", time.Since(now))
}

func exercise(c *catalog, ids ...string) {
	var wg sync.WaitGroup
	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go func(i int) {
			defer wg.Done()
			for _, id := range ids {
				c.get(id)
			}
			c.add("generated_"+strconv.Itoa(i), "generated product")
		}(i + 1)
	}

	wg.Wait()
}

type catalog struct {
	mu   sync.RWMutex
	data map[string]string
}

func (c *catalog) add(id, product string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	// simulate load
	time.Sleep(500 * time.Nanosecond)
	c.data[id] = product
}

func (c *catalog) get(id string) string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	// simulate load
	time.Sleep(500 * time.Nanosecond)
	// avoid key existence check
	return c.data[id]
}

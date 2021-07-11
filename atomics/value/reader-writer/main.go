package main

import (
	"fmt"
	"log"
	"sync"
	"sync/atomic"
)

type Config struct {
	a []int
}

// to test this program, make sure to run it using the -race flag
// go run -race main.go
func main() {
	var wg sync.WaitGroup
	var v atomic.Value
	// to avoid panics when we do type assertion
	v.Store(Config{a: []int{}})

	// writer
	go func() {
		var i int
		for {
			i++
			cfg := Config{
				a: []int{i + 1, i + 2, i + 3, i + 4, i + 5},
			}
			v.Store(cfg)
		}
	}()

	// reader
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			// we're gonna get a panic this way
			// cfg := v.Load().(Config)
			cfg, ok := v.Load().(Config)
			if !ok {
				log.Fatalf("received different type: %T", cfg)
			}
			fmt.Println("cfg", cfg)
		}()
	}
	wg.Wait()
}

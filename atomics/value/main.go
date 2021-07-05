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

func main() {
	var v atomic.Value
	var wg sync.WaitGroup

	// writer
	// we need to wait for the first value
	// to be stored, otherwise the type assertion
	// will fail and panic
	wg.Add(1)
	var once sync.Once
	go func() {
		var i int
		for {
			i++
			cfg := Config{
				a: []int{i + 1, i + 2, i + 3, i + 4, i + 5},
			}
			v.Store(cfg)
			once.Do(func() {
				wg.Done()
			})
		}
	}()

	// reader
	// wait for the first value to be stored
	wg.Wait()
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			// we're gonna get a panic this way
			// cfg := v.Load().(Config)
			cfg, ok := v.Load().(Config)
			if !ok {
				log.Fatalf("received different type: %T", cfg)
			}
			fmt.Println("cfg", cfg)
			wg.Done()
		}()
	}
	wg.Wait()
}

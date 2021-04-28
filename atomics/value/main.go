package main

import (
	"fmt"
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
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			cfg := v.Load()
			fmt.Println("cfg", cfg)
			wg.Done()
		}()
	}
	wg.Wait()
}

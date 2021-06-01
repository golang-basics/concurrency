package main

import (
	"fmt"
	"sync"
)

func main() {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("create object")
			return struct{}{}
		},
	}

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println("go routine", i)
			obj := pool.Get()
			pool.Put(obj)
		}(i)
	}
	wg.Wait()
}

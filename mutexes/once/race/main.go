package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	var wg sync.WaitGroup

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			defer wg.Done()
			once.Do(func() {
				fmt.Println("i:", i)
			})
		}(i)
	}

	wg.Wait()
}

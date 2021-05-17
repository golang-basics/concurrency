package main

import (
	"fmt"
	"sync"
)

func main() {
	namedClosures()
	anonymousClosures()
}

func namedClosures() {
	var wg sync.WaitGroup
	f := func(n int) {
		defer wg.Done()
		fmt.Println("printing from named closure", n)
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go f(i)
	}
	wg.Wait()
}

func anonymousClosures() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Println("printing from anonymoys closure", n)
		}(i)
	}
	wg.Wait()
}

package main

import (
	"fmt"
	"sync"
)

// try running this with the -race flag
// go run -race main.go
func main() {
	var wg sync.WaitGroup
	var count int
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			defer wg.Done()
			// Each go routine only reads the count data and does not expose its data outside
			var localCount int
			localCount += count + i + 1
			fmt.Println("local count", localCount)
		}(i)
	}
	wg.Wait()
}

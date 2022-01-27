package main

import (
	"fmt"
	"sync"
)

// Try running the following program with the -race flag
// go run -race main.go
// the order does not matter at all, the correctness of our concurrent code does
// In the below example we attempt to fetch, increment and update the count variable.
// While everything seems to be correct, our program does not have correctness
// when it comes to running our concurrent code.
// We assume each go routine will have had incremented the count variable when fetching its value,
// which is not the case. This is why we end up with different results. Sometimes we increment
// based on the incremented value, other times we only increment based on the initial value.
func main() {
	var count int
	var wg sync.WaitGroup
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			count++
		}()
	}
	wg.Wait()
	fmt.Println("count", count)
}

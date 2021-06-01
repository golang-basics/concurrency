package main

import (
	"fmt"
	"sync"
)

// to check for race conditions use the -race flag
// go run -race main.go
func main() {
	value := 10

	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		fmt.Println("reading value:", value)
	}()
	go func() {
		defer wg.Done()
		fmt.Println("overwriting value to 15")
		value = 15
	}()
	go func() {
		defer wg.Done()
		fmt.Println()
		fmt.Println("reading then overwriting value to 20")
		fmt.Println("reading value before overwriting", value)
		value = 20
	}()
	wg.Wait()

	fmt.Println("value in main:", value)
}

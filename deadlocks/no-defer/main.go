package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go goodWork(&wg, -1)
	//go badWork(&wg, -1)
	wg.Wait()
}

func goodWork(wg *sync.WaitGroup, n int) {
	defer wg.Done()
	fmt.Println("operation 1")
	if n < 0 {
		fmt.Println("n cannot be zero")
		return
	}
}

func badWork(wg *sync.WaitGroup, n int) {
	fmt.Println("operation 1")
	if n < 0 {
		fmt.Println("n cannot be zero")
		return
	}
	wg.Done()
}

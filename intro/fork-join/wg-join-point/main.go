package main

import (
	"fmt"
	"sync"
	"time"
)

func work() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("printing stuff")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		work()
	}()
	wg.Wait()
	fmt.Println("done waiting, main exits")
}

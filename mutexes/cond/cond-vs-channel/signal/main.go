package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	signal := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		<-signal
		fmt.Println("go routine 1: done")
	}()
	go func() {
		defer wg.Done()
		<-signal
		fmt.Println("go routine 2: done")
	}()

	// wake (signal) 1 waiting go routine
	time.Sleep(500 * time.Millisecond)
	signal <- struct{}{}

	// wake (signal) 1 waiting go routine
	time.Sleep(500 * time.Millisecond)
	signal <- struct{}{}

	wg.Wait()
}

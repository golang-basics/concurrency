package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	cond, broadcast := make(chan struct{}), make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		<-cond
		fmt.Println("go routine 1: done")
	}()
	go func() {
		defer wg.Done()
		<-cond
		fmt.Println("go routine 2: done")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		<-broadcast
		for i := 0; i < 2; i++ {
			cond <- struct{}{}
		}
	}()

	// wake (broadcast) all waiting go routines
	time.Sleep(time.Second)
	broadcast <- struct{}{}

	wg.Wait()
}

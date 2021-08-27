package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	broadcast := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		<-broadcast
		fmt.Println("go routine 1: done")
	}()
	go func() {
		defer wg.Done()
		<-broadcast
		fmt.Println("go routine 2: done")
	}()

	time.Sleep(time.Second)
	close(broadcast)

	wg.Wait()
}

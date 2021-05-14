package main

import (
	"fmt"
	"sync"
)

func main() {
	doneChan()
	doneWaitGroup()
}

func doneChan() {
	done := make(chan struct{})
	go func() {
		fmt.Println("chan: go routine 1 is done")
		done <- struct{}{}
	}()
	go func() {
		fmt.Println("chan: go routine 2 is done")
		done <- struct{}{}
	}()
	for i := 0; i < 2; i++ {
		<-done
	}
	fmt.Println("chan: all go routines are done")
}

func doneWaitGroup() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("wait group: go routine 1 is done")
		wg.Done()
	}()
	go func() {
		fmt.Println("wait group: go routine 2 is done")
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("wait group: all go routines are done")
}

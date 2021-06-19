package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("go routines before work:", runtime.NumGoroutine())
	// conventionally it's called done to signal cancellation of processes
	// also conventionally it's a chan os struct{}, we just need to signal
	// we don't need to pass any kind of memory around go routines
	// also a a convention the done channel is the 1st param, but it does not have to be
	done := make(chan struct{})
	go work(done, nil)
	go func() {
		time.Sleep(time.Second)
		fmt.Println("cancelling all go routines, listening on done")
		close(done)
	}()
	fmt.Println("number of go routines while working:", runtime.NumGoroutine())

	time.Sleep(2 * time.Second)
	fmt.Println("number of go routines after closing done:", runtime.NumGoroutine())
}

func work(done chan struct{}, strings <-chan string) {
	defer fmt.Println("work is done")
	for {
		select {
		case s := <-strings:
			fmt.Println(s)
		case <-done:
			return
		}
	}
}

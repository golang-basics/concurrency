package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// usually 1 go routine goes for main
	// it automatically runs in its own go routine
	fmt.Println("go routines before work:", runtime.NumGoroutine())
	go work(nil)
	timer := time.NewTimer(5 * time.Second)
	for {
		select {
		case <-timer.C:
			fmt.Println("exiting")
			return
		default:
		}
		time.Sleep(time.Second)
		fmt.Println("number of go routines:", runtime.NumGoroutine())
	}
}

func work(strings <-chan string) {
	defer fmt.Println("work is done")
	for s := range strings {
		fmt.Println(s)
	}
}

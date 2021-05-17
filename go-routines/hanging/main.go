package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() {
		// will hang forever
	}()
	// always at least 1 go routine for main
	fmt.Println("number of go routines", runtime.NumGoroutine())
	time.Sleep(time.Second)
	fmt.Println("main is done")
}

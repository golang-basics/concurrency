package main

import (
	"fmt"
	"runtime"
)

// try running the example on Go version < 1.14
func main() {
	runtime.GOMAXPROCS(1)
	go fmt.Println("I try to print")
	// enable below to statement to allow main go routine
	// to be preempted so that other go routines can take execution
	// runtime.Gosched()
	for {
	}
}

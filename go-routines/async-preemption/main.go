package main

import "runtime"

// Asynchronous Preemption Go >= 1.14
func main() {
	runtime.GOMAXPROCS(1)
	go println("goroutine ran")
	for {
	}
}

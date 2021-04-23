package main

import "runtime"

func main() {
	runtime.GOMAXPROCS(1)
	go println("goroutine ran")
	for {
	}
}

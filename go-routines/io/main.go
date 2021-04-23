package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	go fmt.Println("I will print in a different thread")
	// the call below is a syscall, meaning it is a synchronous call
	// a separate thread will take care of it to avoid blocking the existing thread
	_, _ = os.Create("test.txt")
}

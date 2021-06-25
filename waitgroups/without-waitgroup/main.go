package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// Go Scheduler will create 16 threads
	// which in theory will be able to execute
	// 16 blocking go routines in parallel
	// totalling ~100ms
	fmt.Println("number of cores:", runtime.NumCPU())
	for i := 0; i < 10; i++ {
		go work(i + 1)
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Println("main is done")
}

func work(id int) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("task", id, "is done")
}

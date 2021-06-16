package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func main() {
	fmt.Println("number of processors(CPUS)", runtime.NumCPU())
	fmt.Println("PID", os.Getpid())
	// why does it not work?
	runtime.GOMAXPROCS(32)
	time.Sleep(time.Minute)
}

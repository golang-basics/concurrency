package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func main() {
	fmt.Println("number of cores", runtime.NumCPU())
	fmt.Println("PID", os.Getpid())
	runtime.GOMAXPROCS(32)
	time.Sleep(time.Minute)
}

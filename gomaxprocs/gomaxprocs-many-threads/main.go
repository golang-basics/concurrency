package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func main() {
	fmt.Println("PID", os.Getpid())
	runtime.GOMAXPROCS(32)
	for i := 0; i < 32; i++ {
		go func() {
			for {
			}
		}()
	}
	time.Sleep(time.Minute)
}

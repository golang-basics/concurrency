package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

// go build
// GOMAXPROCS=1 GOGC=off GODEBUG=schedtrace=1000,scheddetail=1 ./io
func main() {
	// keep in mind the values from:
	// P0: schedtick
	// P0: syscalltick
	// https://github.com/golang/go/blob/master/src/runtime/runtime2.go#L608
	// https://github.com/golang/go/blob/master/src/runtime/runtime2.go#L609
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second)
		fmt.Println("go routine 1: done")
	}()
	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Second)
		fmt.Println("go routine 2: done")
	}()
	go func() {
		defer wg.Done()
		file, _ := os.CreateTemp("", "test.txt")
		// 1GB write
		_, _ = file.Write([]byte(strings.Repeat("a", 1_000_000_000)))
		fmt.Println("done writing")
	}()
	wg.Wait()

	// wait for one more Tracing event
	time.Sleep(2 * time.Second)
}

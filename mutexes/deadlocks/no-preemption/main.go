package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// GOROOT/src/runtime/proc.go#5206
// https://github.com/golang/go/blob/35ea62468bf7e3a79011c3ad713e847daa9a45a2/src/runtime/proc.go#L4159-L4233
// for{} loops in conjunction with runtime.GOMAXPROCS(1) will make a go routine non-preemptive, or hanging forever
func main() {
	runtime.GOMAXPROCS(1)
	var mu sync.Mutex
	var count int
	go func() {
		mu.Lock()
		defer mu.Unlock()
		count = 10
		// No longer an issue for Go 1.14+
		for {
		}
	}()

	time.Sleep(time.Second)
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Println("count", count)
}

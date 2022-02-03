package main

import (
	"fmt"
	"sync"
	"time"
)

// EXERCISE:
// Find what's wrong in the code below
// Make sure all the tests are passing
// DO NOT remove/modify any Sleep() calls
// DO NOT remove/modify any read()/write() calls

// try running this with the -race flag
// go run -race exercise.go

// run the tests using:
// GOFLAGS="-count=1" go test .

func main() {
	// p1 and p2 are the number of executions per process
	p1, p2 := exercise(time.Second)
	fmt.Println("p1:", p1)
	fmt.Println("p2:", p2)
}

func exercise(d time.Duration) (int, int) {
	var p1, p2 int
	var f file

	for begin := time.Now(); time.Since(begin) < d; {
		go func() {
			f.write() // 3ns
			p1++
		}()
		go func() {
			f.read() // 1ns
			f.read() // 1ns
			f.read() // 1ns
			p2++
		}()
	}

	return p1, p2
}

type file struct {
	mu sync.Mutex
}

func (f *file) read() {
	f.mu.Lock()
	defer f.mu.Unlock()
	time.Sleep(time.Nanosecond)
}

func (f *file) write() {
	f.mu.Lock()
	defer f.mu.Unlock()
	time.Sleep(3 * time.Nanosecond)
}

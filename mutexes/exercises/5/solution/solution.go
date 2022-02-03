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

// SOLUTION
// The problem we're solving here is Starvation.
// Even if all operations are equal in terms of time burst.
// Because we have 3 calls to read and 1 call to write,
// the mutex local to each call is being used more frequently,
// thus allowing one process to acquire it way more often than the other.
// The fix simply using the mutex evenly distributing it across workers/Gs.
// In other words, just use the mutex directly, not inside the read call.
// Always be careful with mutexes local to the methods and test for starvation.
func main() {
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
			f.mu.Lock()
			f.read() // 1ns
			f.read() // 1ns
			f.read() // 1ns
			f.mu.Unlock()
			p2++
		}()
	}

	return p1, p2
}

type file struct {
	mu sync.Mutex
}

func (f *file) read() {
	time.Sleep(time.Nanosecond)
}

func (f *file) write() {
	f.mu.Lock()
	defer f.mu.Unlock()
	time.Sleep(3 * time.Nanosecond)
}

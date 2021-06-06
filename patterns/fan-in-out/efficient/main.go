//go:generate go run ../../codegen/main.go -tpl=fanin -type=int
//go:generate go run ../../codegen/main.go -tpl=repeatfn -type=int
//go:generate go run ../../codegen/main.go -tpl=take -type=int
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

// before running main.go, make sure to run:
// go generate main.go
func main() {
	done := make(chan struct{})
	defer close(done)
	random := func() int {
		return rand.Intn(50000000)
	}
	start := time.Now()
	randIntStream := RepeatFn(done, random)
	numFinders := runtime.NumCPU()
	finders := make([]<-chan int, numFinders)
	for i := 0; i < numFinders; i++ {
		finders[i] = primeFinder(done, randIntStream)
	}

	fmt.Println("primes:")
	for prime := range Take(done, FanIn(done, finders...), 10) {
		fmt.Println("prime:", prime)
	}
	fmt.Printf("search took: %v", time.Since(start))
}

func primeFinder(done <-chan struct{}, intStream <-chan int) <-chan int {
	primeStream := make(chan int)
	go func() {
		defer close(primeStream)
		for integer := range intStream {
			integer -= 1
			prime := true
			for divisor := integer - 1; divisor > 1; divisor-- {
				if integer%divisor == 0 {
					prime = false
					break
				}
			}

			if prime {
				select {
				case <-done:
					return
				case primeStream <- integer:
				}
			}
		}
	}()
	return primeStream
}

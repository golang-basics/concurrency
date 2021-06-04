package main

import (
	"fmt"
	"math/rand"
	"time"

	"concurrency/patterns/generators"
)

var (
	take     = generators.Take
	toInt    = generators.ToInt
	repeatFn = generators.RepeatFn
)

func main() {
	done := make(chan struct{})
	defer close(done)
	random := func() interface{} {
		return rand.Intn(50000000)
	}
	start := time.Now()

	randIntStream := toInt(done, repeatFn(done, random))
	fmt.Println("primes:")
	for prime := range take(done, primeFinder(done, randIntStream), 10) {
		fmt.Println("prime:", prime)
	}
	fmt.Printf("search took: %v", time.Since(start))
}

func primeFinder(done <-chan struct{}, intStream <-chan int) <-chan interface{} {
	primeStream := make(chan interface{})
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

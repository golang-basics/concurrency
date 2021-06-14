package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})
	defer close(done)
	interval := time.Second
	requests := newRequests(done, 5)
	limiter := newLimiter(done, interval, 1)
	fmt.Println("BURST: 1")
	for req := range requests {
		<-limiter
		fmt.Println("request", req, "arrived at", time.Now())
	}

	burstyRequests := newRequests(done, 5)
	burstyLimiter := newLimiter(done, interval, 3)
	fmt.Println("BURST: 3")
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, "arrived at", time.Now())
	}
}

func newRequests(done chan struct{}, buffer int) <-chan int {
	out := make(chan int, buffer)
	defer close(out)
	for i := 0; i < buffer; i++ {
		select {
		case <-done:
			return out
		case out <- i + 1:
		}
	}
	return out
}

func newLimiter(done chan struct{}, interval time.Duration, burst int) <-chan time.Time {
	limiter := make(chan time.Time, burst)
	if burst == 1 {
		return time.Tick(interval)
	}

	for i := 0; i < burst; i++ {
		select {
		case <-done:
			return limiter
		case limiter <- time.Now():
		}
	}
	go func() {
		for t := range time.Tick(interval) {
			select {
			case <-done:
				return
			case limiter <- t:
			}
		}
	}()

	return limiter
}

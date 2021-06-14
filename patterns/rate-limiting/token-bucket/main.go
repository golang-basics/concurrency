package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	done := make(chan struct{})
	defer close(done)

	bucket := newTokenBucket(done, time.Second, 3)

	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			defer wg.Done()
			bucket.Wait()
			fmt.Println("request", i)
		}(i + 1)
	}

	wg.Wait()
}

func newTokenBucket(done chan struct{}, r time.Duration, b int) tokenBucket {
	tokens := make(chan time.Time, b)
	for i := 0; i < b; i++ {
		tokens <- time.Now()
	}
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-time.Tick(r):
				tokens <- t
			}
		}
	}()
	return tokenBucket{
		rate:   r,
		tokens: tokens,
	}
}

type tokenBucket struct {
	rate   time.Duration
	tokens chan time.Time
}

func (t tokenBucket) Wait() {
	select {
	case <-t.tokens:
	}
}

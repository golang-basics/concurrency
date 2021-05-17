package main

import (
	"fmt"
	"time"
)

type request func()

func main() {
	requests := make([]request, 0)
	for i := 1; i <= 100; i++ {
		f := func() {
			time.Sleep(200 * time.Millisecond)
			// the value will be accessed from heap
			// if it has not been evaluated and the function returned
			fmt.Println("request", i)
		}
		requests = append(requests, f)
	}

	for _, r := range requests {
		r()
	}
}

package main

import (
	"fmt"

	"concurrency/patterns/cancellation"
)

var (
	inc = cancellation.Inc
	dec = cancellation.Dec
	sq  = cancellation.Sq
)

func main() {
	done := make(chan struct{})
	defer close(done)
	nums := cancellation.Gen(done, 1, 2, 3)
	for v := range dec(done, sq(done, inc(done, nums))) {
		fmt.Println("value:", v)
	}
}

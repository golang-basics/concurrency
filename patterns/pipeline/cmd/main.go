package main

import (
	"fmt"

	"concurrency/patterns/pipeline"
)

var (
	gen = pipeline.Gen
	inc = pipeline.Inc
	dec = pipeline.Dec
	sq  = pipeline.Sq
)

func main() {
	done := make(chan struct{})
	defer close(done)

	// 1, 2, 3
	nums := gen(done, 1, 2, 3)
	// 2, 3, 4
	incremented := inc(done, nums)
	// 4, 9, 16
	squared := sq(done, incremented)
	// 3, 8, 15
	res := dec(done, squared)
	for n := range res {
		fmt.Println(n)
	}

	fmt.Println("the same exact result using nested calls")
	for n := range dec(done, sq(done, inc(done, gen(done, 1, 2, 3)))) {
		fmt.Println(n)
	}

	fmt.Println("the same exact result using chaining")
	for n := range pipeline.New(1, 2, 3).Increment().Square().Decrement().Result() {
		fmt.Println(n)
	}
}

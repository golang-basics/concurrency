package main

import (
	"fmt"

	"concurrency/patterns/fanout"
	"concurrency/patterns/generator"
)

func main() {
	done := make(chan struct{})
	defer close(done)

	odd := generator.OddIntGen(10)
	c1 := fanout.FanOut(done, odd)
	c2 := fanout.FanOut(done, odd)
	c3 := fanout.FanOut(done, odd)
	c4 := fanout.FanOut(done, odd)

	display(c1, "c1")
	display(c2, "c2")
	display(c3, "c3")
	display(c4, "c4")
}

func display(in <-chan int, label string) {
	for v := range in {
		fmt.Printf("value %s: %d\n", label, v)
	}
}

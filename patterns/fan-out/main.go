package main

import (
	"fmt"

	patterns "concurrency/patterns/fan-out/pkg"
	generators "concurrency/patterns/generator/pkg"
)

func main() {
	odd := generators.OddIntGen(10)
	c1 := patterns.FanOut(odd)
	c2 := patterns.FanOut(odd)
	c3 := patterns.FanOut(odd)
	c4 := patterns.FanOut(odd)
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

package main

import (
	"fmt"

	"concurrency/patterns/fanin"
	"concurrency/patterns/fanout"
	"concurrency/patterns/generator"
)

var (
	intGen = generator.OddIntGen
	fanOut = fanout.FanOut
	fanIn  = fanin.FanIn
)

func main() {
	done := make(chan struct{})
	defer close(done)

	odd := intGen(20)
	c1 := fanOut(done, odd)
	c2 := fanOut(done, odd)
	c3 := fanOut(done, odd)

	for v := range fanIn(done, c1, c2, c3) {
		fmt.Println("value:", v)
	}
}

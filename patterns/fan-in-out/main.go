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
	odd := intGen(20)
	c1, c2, c3 := fanOut(odd), fanOut(odd), fanOut(odd)
	for v := range fanIn(c1, c2, c3) {
		fmt.Println("value:", v)
	}
}

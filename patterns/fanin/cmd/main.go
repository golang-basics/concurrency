package main

import (
	"fmt"

	"concurrency/patterns/fanin"
	"concurrency/patterns/generator"
)

func main() {
	done := make(chan struct{})
	defer close(done)

	odd := generator.OddIntGen(5)
	even := generator.EvenIntGen(5)
	hex := generator.HexIntGen(5)
	out := fanin.FanIn(done, odd, even, hex)

	for n := range out {
		fmt.Println("fanned number:", n)
	}
}

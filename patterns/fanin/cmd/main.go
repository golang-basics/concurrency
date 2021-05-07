package main

import (
	"fmt"

	"concurrency/patterns/fanin"
	"concurrency/patterns/generator"
)

func main() {
	odd := generator.OddIntGen(5)
	even := generator.EvenIntGen(5)
	hex := generator.HexIntGen(5)
	out := fanin.FanIn(odd, even, hex)
	for n := range out {
		fmt.Println("fanned number:", n)
	}
}

package main

import (
	"fmt"

	pattern "concurrency/patterns/fan-in/pkg"
	generators "concurrency/patterns/generator/pkg"
)

func main() {
	odd := generators.OddIntGen(5)
	even := generators.EvenIntGen(5)
	hex := generators.HexIntGen(5)
	out := pattern.FanIn(odd, even, hex)
	for n := range out {
		fmt.Println("fanned number:", n)
	}
}

// The generator pattern is pretty simple and it resembles that
// there is a function which generates a channel of values
// and returns it at the end, while values are pushed by go routine(s)
// the channel must eventually close to avoid dead locks

package main

import (
	"fmt"

	"concurrency/patterns/generator"
)

func main() {
	for evenInt := range generator.EvenIntGen(5) {
		fmt.Println("even int:", evenInt)
	}
	for oddInt := range generator.OddIntGen(5) {
		fmt.Println("odd int:", oddInt)
	}
	for hexInt := range generator.HexIntGen(5) {
		fmt.Println("hex int:", hexInt)
	}
	for word := range generator.WordGen(5) {
		fmt.Println("word:", word)
	}
}

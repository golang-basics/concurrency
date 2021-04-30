// The generator pattern is pretty simple and it resembles that
// there is a function which generates a channel of values
// and returns it at the end, while values are pushed by go routine(s)
// the channel must eventually close to avoid dead locks

package main

import (
	"fmt"

	"concurrency/patterns/generator/pkg"
)

func main() {
	for evenInt := range pkg.EvenIntGen(5) {
		fmt.Println("even int:", evenInt)
	}
	for oddInt := range pkg.OddIntGen(5) {
		fmt.Println("odd int:", oddInt)
	}
	for hexInt := range pkg.HexIntGen(5) {
		fmt.Println("hex int:", hexInt)
	}
	for word := range pkg.WordGen(5) {
		fmt.Println("word:", word)
	}
}

package main

import (
	"fmt"
	"math/rand"

	"concurrency/patterns/generators"
)

func main() {
	done := make(chan struct{})
	defer close(done)

	randFn := func() interface{} { return rand.Int() }
	for num := range generators.Take(done, generators.RepeatFn(done, randFn), 10) {
		fmt.Println(num)
	}
}

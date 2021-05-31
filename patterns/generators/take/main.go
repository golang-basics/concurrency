package main

import (
	"fmt"

	"concurrency/patterns/generators"
)

func main() {
	done := make(chan struct{})
	defer close(done)

	for num := range generators.Take(done, generators.Repeat(done, 1), 10) {
		fmt.Printf("%v ", num)
	}
}

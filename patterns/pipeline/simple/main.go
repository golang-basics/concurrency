package main

import (
	"fmt"
)

// Here are couple of simple rules regarding pipelines:
// 1. They receive some data apply an operation (stage) and return the data
// 2. Usually the output of 1 pipeline is the input of another pipeline
// 3. Thus, each stage in a pipeline operates with the same type of data that it receives/returns
// 4. Pipelines can implemented as a batch (all at once) or as a stream (1 at a time)
func main() {
	numbers := gen(4)
	res := add(multiply(numbers, 2), 1)
	for _, num := range multiply(res, 2) {
		fmt.Println("number:", num)
	}
}

func gen(n int) []int {
	out := make([]int, n)
	for i := 0; i < n; i++ {
		out[i] = i + 1
	}
	return out
}

// add receives a slice of int adds the additive to each element
// and returns a new slice of int
func add(numbers []int, additive int) []int {
	out := make([]int, len(numbers))
	for i, num := range numbers {
		out[i] = num + additive
	}
	return out
}

// multiply receives a slice of int multiplies each element with multiplier
// and returns a new slice of int
func multiply(numbers []int, multiplier int) []int {
	out := make([]int, len(numbers))
	for i, num := range numbers {
		out[i] = num * multiplier
	}
	return out
}

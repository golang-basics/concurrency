package main

import (
	"fmt"
)

func main() {
	numbers := gen(4)
	// Notice here we don't create new slices on each pipeline stage.
	// Instead we use stream pipeline stages that produce only 1 value at a time,
	// thus keeping the memory footprint lower
	fmt.Println("stream")
	for _, num := range numbers {
		res := sMultiply(sAdd(sMultiply(num, 2), 1), 2)
		fmt.Println("number:", res)
	}

	// Here we create a brand new slice on each pipeline stage
	fmt.Println("batch")
	for _, num := range bMultiply(bAdd(bMultiply(numbers, 2), 1), 2) {
		fmt.Println("number:", num)
	}
}

func sAdd(value int, additive int) int {
	return value + additive
}

func sMultiply(value int, multiplier int) int {
	return value * multiplier
}

func bAdd(numbers []int, additive int) []int {
	out := make([]int, len(numbers))
	for i, num := range numbers {
		out[i] = num + additive
	}
	return out
}

func bMultiply(numbers []int, multiplier int) []int {
	out := make([]int, len(numbers))
	for i, num := range numbers {
		out[i] = num * multiplier
	}
	return out
}

func gen(n int) []int {
	out := make([]int, n)
	for i := 0; i < n; i++ {
		out[i] = i + 1
	}
	return out
}

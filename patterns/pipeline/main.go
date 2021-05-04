package main

import (
	"fmt"

	pipeline "concurrency/patterns/pipeline/pkg"
)

func main() {
	// 1, 2, 3
	nums := gen(1,2,3)
	// 2, 3, 4
	incremented := inc(nums)
	// 4, 9, 16
	squared := sq(incremented)
	// 3, 8, 15
	res := dec(squared)
	for n := range res {
		fmt.Println(n)
	}

	fmt.Println("the same exact result using nested calls")
	for n := range dec(sq(inc(gen(1,2,3)))) {
		fmt.Println(n)
	}

	fmt.Println("the same exact result using chaining")
	for n := range pipeline.New(1,2,3).Increment().Square().Decrement().Result() {
		fmt.Println(n)
	}
}

func gen(vs ...int) chan int {
	out := make(chan int)
	go func() {
		for _, n := range vs {
			out <- n
		}
		close(out)
	}()
	return out
}

func inc(in <-chan int) chan int {
	out := make(chan int)
	go func() {
		for i := range in {
			out <- i + 1
		}
		close(out)
	}()
	return out
}

func dec(in <-chan int) chan int {
	out := make(chan int)
	go func() {
		for i := range in {
			out <- i - 1
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) chan int {
	out := make(chan int)
	go func() {
		for i := range in {
			out <- i * i
		}
		close(out)
	}()
	return out
}

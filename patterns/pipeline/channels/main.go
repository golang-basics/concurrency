package main

import "fmt"

func main() {
	done := make(chan struct{})
	defer close(done)

	genStage := gen(done, 1, 2, 3, 4)
	multiplyStage := multiply(done, genStage, 2)
	addStage := add(done, multiplyStage, 1)
	res := multiply(done, addStage, 2)

	for num := range res {
		fmt.Println("number:", num)
	}
}

func gen(done chan struct{}, numbers ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, num := range numbers {
			select {
			case <-done:
				return
			case out <- num:
			}
		}
	}()
	return out
}

func add(done chan struct{}, in <-chan int, additive int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			select {
			case <-done:
				return
			case out <- num + additive:
			}
		}
	}()
	return out
}

func multiply(done chan struct{}, in <-chan int, multiplier int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			select {
			case <-done:
				return
			case out <- num * multiplier:
			}
		}
	}()
	return out
}

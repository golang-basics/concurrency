package main

import "fmt"

func pass(left, right chan func(), i int) {
	fn := <-right
	left <- func() {
		fn()
		fmt.Println("gopher:", i)
	}
}

func main() {
	leftmost := make(chan func())
	left := leftmost
	right := leftmost

	for i := 0; i < 10; i++ {
		right = make(chan func())
		go pass(left, right, i+1)
		left = right
	}

	right <- func() {
		fmt.Println("initial function")
	}
	fn := <-leftmost
	fn()
}

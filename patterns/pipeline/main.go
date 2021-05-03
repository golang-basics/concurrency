package main

import "fmt"

func main() {
	nums := gen(3) // 1, 2, 3
	inc := inc(nums) // 2, 3, 4
	sq := sq(inc) // 4, 9, 16
	res := dec(sq) // 3, 8, 15
	for n := range res {
		fmt.Println(n)
	}
}

func gen(n int) chan int {
	out := make(chan int)
	go func() {
		for i := 1; i <= n; i++ {
			out <- i
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

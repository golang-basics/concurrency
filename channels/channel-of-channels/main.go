package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{1, 2, 3}
	for v := range bridge(channels(inc(nums), dec(nums), sq(nums))) {
		fmt.Println("value:", v)
	}
}

func channels(cs ...chan int) chan chan int {
	out := make(chan chan int, len(cs))
	for _, ch := range cs {
		out <- ch
	}
	close(out)
	return out
}

func bridge(in chan chan int) chan int {
	out := make(chan int)
	go func() {
		var wg sync.WaitGroup
		multiplex := func(c chan int) {
			for v := range c {
				out <- v
			}
			wg.Done()
		}
		for c := range in {
			wg.Add(1)
			go multiplex(c)
		}
		wg.Wait()
		close(out)
	}()
	return out
}

func inc(numbers []int) chan int {
	out := make(chan int)
	go func() {
		for _, i := range numbers {
			out <- i + 1
		}
		close(out)
	}()
	return out
}

func dec(numbers []int) chan int {
	out := make(chan int)
	go func() {
		for _, i := range numbers {
			out <- i - 1
		}
		close(out)
	}()
	return out
}

func sq(numbers []int) chan int {
	out := make(chan int)
	go func() {
		for _, i := range numbers {
			out <- i * i
		}
		close(out)
	}()
	return out
}

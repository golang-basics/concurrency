package main

import "fmt"

func main() {
	ch := make(chan int)
	go write(ch)
	for i := range ch {
		fmt.Println(i)
	}
}

func write(ch chan int) {
	for i := 0; i < 3; i++ {
		ch <- i
	}
	close(ch)
}

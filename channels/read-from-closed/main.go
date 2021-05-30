package main

import "fmt"

func main() {
	intChan := make(chan int)
	close(intChan)

	value, ok := <-intChan
	fmt.Println("value:", value)
	fmt.Println("ok:", ok)

	value, ok = <-intChan
	fmt.Println("value:", value)
	fmt.Println("ok:", ok)

	value, ok = <-intChan
	fmt.Println("value:", value)
	fmt.Println("ok:", ok)
}

package main

import "fmt"

func main() {
	ch := make(chan bool, 1)
	ch <- true
	val, ok := <-ch
	fmt.Println("val", val)
	fmt.Println("ok", ok)

	close(ch)
	val, ok = <-ch
	fmt.Println("val", val)
	fmt.Println("ok", ok)
}

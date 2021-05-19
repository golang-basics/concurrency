package main

import (
	"fmt"
	"unsafe"
)

func main() {
	emptyStructChan := make(chan struct{})
	intChan := make(chan int)
	go func() {
		emptyStructChan <- struct{}{}
		intChan <- 1
	}()
	fmt.Println(unsafe.Sizeof(<-emptyStructChan)) // 0
	fmt.Println(unsafe.Sizeof(<-intChan))         // 8 => 8*8=64 => int64
}

package main

import (
	"fmt"
	"time"
)

func main() {
	p1, p2 := make(chan struct{}), make(chan struct{})
	go func() {
		// blocked forever, waiting on p2
		<-p2
		p1 <- struct{}{}
		fmt.Println("done in go routine")
	}()
	// blocked forever, waiting on p1
	<-p1
	p2 <- struct{}{}
	fmt.Println("done in main")
	time.Sleep(500*time.Millisecond)
}

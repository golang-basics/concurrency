package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	begin := make(chan struct{})
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-begin
			fmt.Println("go routine", i)
		}(i + 1)
	}

	fmt.Println("unblocking 1 go routine")
	begin <- struct{}{}

	time.Sleep(time.Second)
	fmt.Println("unblocking the rest of the go routines")
	close(begin)
	wg.Wait()
}

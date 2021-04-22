package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(100 * time.Millisecond)
	timer := time.NewTimer(3 * time.Second)

	for {
		select {
		case t := <-ticker.C:
			// usually check for something once in a while
			// or try an action till it succeeds or the select times out
			fmt.Println("tick", t)
		case <-timer.C:
			fmt.Println("time is up")
			return
		}
	}
}

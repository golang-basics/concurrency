package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.NewTicker(100 * time.Millisecond)
	timeout := time.NewTimer(3 * time.Second)
	for {
		select {
		case t := <-tick.C:
			fmt.Println("tick", t)
		case <-timeout.C:
			fmt.Println("timeout")
			return
		}
	}
}

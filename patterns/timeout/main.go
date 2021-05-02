package main

import (
	"fmt"
	"time"
)

func main() {
	timeout := time.NewTimer(3 * time.Second)
	for {
		time.Sleep(500 * time.Millisecond)
		select {
		case <-timeout.C:
			fmt.Println("timeout")
			return
		default:
			fmt.Println("retrying")
		}
	}
}

package main

import (
	"fmt"
	"time"
)

// sometimes we never know when and if a channel is ready
// a good practice is to always have a fallback
// a fallback can be a timeout or a default case
// to prevent the select statement from blocking
func main() {
	done := make(chan struct{})
	select {
	case <-done:
		fmt.Println("done")
	case <-time.After(time.Second):
		fmt.Println("timed out")
	}
}

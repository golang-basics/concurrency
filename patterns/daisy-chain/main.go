package main

import (
	"fmt"

	"concurrency/patterns/cancellation"
)

func main() {
	done := make(chan struct{})
	defer close(done)
	nums := cancellation.Gen(done, 1, 2, 3)
	for v := range cancellation.Inc(done, nums) {
		fmt.Println("value:", v)
	}
}

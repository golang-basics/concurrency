// The Daisy-Chain pattern is like a chain reaction or Domino game.
// Or to better illustrate this from Electrical Engineering,
// think of Daisy-Chains as a long LED strip, where the current flows
// through each LED, powering the LED and leading the current further to the next LED

// Here's how the Daisy-Chain flow looks like:
// 1. We arrange all the pieces on the table, by spinning some go routines
// 2. On each iteration we reassign left to previous right channel
// 3. When we call the pass function results propagate from the rightmost (last) channel
// till the leftmost (first) channel
// 4. All that's left is to trigger the chain by pushing a value into the rightmost channel
// by which all go routines are blocked and waiting for
// Once that's done, the chain reaction takes place, hence the name of Daisy-Chain.
// For more about the Daisy-Chain pattern check out the wiki:
// https://en.wikipedia.org/wiki/Daisy_chain_(electrical_engineering)

package main

import (
	"bytes"
	"fmt"
)

func pass(left, right chan int, buff *bytes.Buffer) {
	value := <-right
	buff.WriteString(fmt.Sprintf("%d <-- ", value))
	left <- 1 + value
}

func main() {
	leftmost := make(chan int)
	left := leftmost
	right := leftmost

	var buff bytes.Buffer
	for i := 0; i < 100; i++ {
		right = make(chan int)
		go pass(left, right, &buff)
		left = right
	}

	right <- 1
	buff.WriteString(fmt.Sprintf("%d", <-leftmost))
	fmt.Println(buff.String())
}

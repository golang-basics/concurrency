// The ping pong pattern is very simple and useful in a lot of cases.
// It basically does a ping pong in terms of concurrency.
// The data is ping-ponged between multiple go routines back an forth
// There's constantly some go routine receiving data (PING) and re-passing (PONG)
// to another go routine listening on the same data.
// Huge thanks to divan.dev. Check out the full resource here: https://divan.dev/posts/go_concurrency_visualize/

package main

import (
	"fmt"
	"time"
)

func main() {
	table := make(chan int)
	go player(table, 1)
	go player(table, 2)
	go player(table, 3)

	table <- 0
	time.Sleep(1 * time.Second)
	fmt.Println(<-table)
}

func player(table chan int, playerNo int) {
	for {
		ball := <-table
		fmt.Println("got ball:", ball, "from player:", playerNo)
		ball++
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}

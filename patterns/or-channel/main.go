// The Or channel pattern can be very useful when selecting
// between multiple done actions and only 1 done (fastest) is important.
// The problem can be solved using a good old select statement,
// however it can be more convenient to just have a 1 liner, which does the job.
// You can find more about the Or channel pattern on:
// https://www.oreilly.com/library/view/concurrency-in-go/9781491941294/ch04.html

package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	<-or(
		operation(2*time.Hour, "1"),
		operation(5*time.Minute, "2"),
		operation(1*time.Second, "3"),
		operation(1*time.Hour, "4"),
		operation(1*time.Minute, "5"),
	)
	fmt.Printf("done after %v", time.Since(start))
}

func or(channels ...<-chan interface{}) <-chan interface{} {
	// recursive exit condition
	switch len(channels) {
	case 0:
		// no channels passed
		return nil
	case 1:
		// return the first channel passed
		return channels[0]
	}

	orDone := make(chan interface{})
	go func() {
		defer close(orDone)
		switch len(channels) {
		case 2:
			// try to select the channel which closes first
			// small optimization for 2 channels
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			// try to select the channel which closes first
			// remember select does not fall through case by case
			// it tries to resolve all cases simultaneously picking the one who's most ready
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			// here happens the recursive call => channels > 3
			// the orDone channel is appended on each recursive call
			// so that the rest of go routines can exit
			// when the first go routine has exited and orDone was closed
			case <-or(append(channels[3:], orDone)...):
			}
		}
	}()
	return orDone
}

func operation(after time.Duration, name string) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
		fmt.Println("operation:", name, "executed")
	}()
	return c
}

// The same exact example as the one from "simple" directory
// Except we don't close the outbound channels
// And only executing 2 iterations (which will return before the done timeout)
package main

import (
	"fmt"
	"time"
)

func main() {
	timeout := 2 * time.Second
	done := make(chan struct{})
	time.AfterFunc(10*time.Second, func() {
		close(done)
	})

	heartbeat, results := work(done, timeout/2)
	for {
		select {
		case _, ok := <-heartbeat:
			if !ok {
				return
			}
			fmt.Println("pulse")
		case r, ok := <-results:
			if !ok {
				return
			}
			fmt.Println("result", r.Second())
		case <-time.After(timeout):
			fmt.Println("worker go routine is not healthy")
			return
		}
	}
}

func work(done chan struct{}, interval time.Duration) (chan struct{}, chan time.Time) {
	heartbeat := make(chan struct{})
	results := make(chan time.Time)
	go func() {
		pulse := time.Tick(interval)
		work := time.Tick(interval * 2)
		sendPulse := func() {
			select {
			case heartbeat <- struct{}{}:
			default:
			}
		}
		sendResult := func(r time.Time) {
			for {
				select {
				case <-done:
					return
				case <-pulse:
					sendPulse()
				case results <- r:
					return
				}
			}
		}

		// we try to return earlier than the outside timeout
		// closes the done channel
		// also we don't close the channels before we return
		for i := 0; i < 2; i++ {
			select {
			case <-done:
				return
			case <-pulse:
				sendPulse()
			case r := <-work:
				sendResult(r)
			}
		}
	}()
	return heartbeat, results
}

// The HeartBeat pattern is very much like the human heart beat.
// As the human heart beats at a certain rate/interval while
// the entire body does a multitude of functions, the same goes
// for any concurrent process that does some amount of work.
// The pulse of beats at a certain rate/interval to tell the
// processes outside the concurrent process (go routine) that it's alive
// and it's doing some work. This can be very helpful to diagnose or
// monitor running/long lived concurrent processes like: workers, daemons and so on.
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
			return
		}
	}
}

// work produces two types of values: a [heartbeat value] and a [time value].
// the consumer can read both values or simply use the results value
// if diagnosing the running process is not a concern.
func work(done chan struct{}, interval time.Duration) (chan struct{}, chan time.Time) {
	heartbeat := make(chan struct{})
	results := make(chan time.Time)
	go func() {
		defer close(heartbeat)
		defer close(results)

		// here we try to simulate work being done
		// slower than the actual pulse of the heartbeat
		pulse := time.Tick(interval)
		work := time.Tick(interval * 2)
		sendPulse := func() {
			select {
			case heartbeat <- struct{}{}:
			default:
				// the default is very important so that
				// the send on heartbeat does not block
				// in case nobody is interested in reading
				// from the heartbeat
			}
		}
		sendResult := func(r time.Time) {
			for {
				select {
				case <-done:
					return
				case <-pulse:
					// send pulse on channel send (results <- r)
					sendPulse()
				case results <- r:
					return
				}
			}
		}

		for {
			select {
			case <-done:
				return
			case <-pulse:
				// send pulse on channel receive (r := <-work)
				sendPulse()
			case r := <-work:
				sendResult(r)
			}
		}
	}()
	return heartbeat, results
}

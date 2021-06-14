package main

import (
	"log"
	"os"
	"time"
)

func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)

	done := make(chan struct{})
	time.AfterFunc(11*time.Second, func() {
		close(done)
	})

	steward := newSteward(4*time.Second, ward)
	for range steward(done, 4*time.Second) {
	}
	log.Println("done")
}

// startGoroutineFn represents the function that can be monitored
type startGoroutineFn func(done <-chan struct{}, interval time.Duration) (heartbeat <-chan struct{})

// ward represents really any concurrent process that will be monitored
// by a steward (supervisor) go routine.
func ward(done <-chan struct{}, interval time.Duration) <-chan struct{} {
	heartbeat := make(chan struct{})
	log.Println("ward: hello")
	pulse := time.Tick(interval)
	var i int
	go func() {
		for {
			if i == 2 {
				// break things on purpose
				// to simulate some kind of
				// death spiral
				log.Println("ward: death spiral")
				pulse = nil
			}
			select {
			case <-done:
				log.Println("ward: halting")
				return
			case <-pulse:
				log.Println("ward: pulse")
				select {
				case heartbeat <- struct{}{}:
					i++
				default:
				}
			}
		}
	}()
	return heartbeat
}

// newSteward runs a go routine that monitors other go routines.
// Notice, it also returns startGoroutineFn which itself can be monitored
func newSteward(timeout time.Duration, startGoroutine startGoroutineFn) startGoroutineFn {
	return func(done <-chan struct{}, interval time.Duration) <-chan struct{} {
		heartbeat := make(chan struct{})
		go func() {
			defer close(heartbeat)

			var wardDone chan struct{}
			var wardHeartbeat <-chan struct{}
			startWard := func() {
				wardDone = make(chan struct{})
				wardHeartbeat = startGoroutine(or(wardDone, done), timeout/2)
			}

			startWard()
			pulse := time.Tick(interval)

		loop:
			for {
				timer := time.After(timeout)

				for {
					select {
					case <-pulse:
						log.Println("steward: pulse")
						select {
						case heartbeat <- struct{}{}:
						default:
						}
					case <-wardHeartbeat:
						continue loop
					case <-timer:
						log.Println("steward: ward is unhealthy; restarting")
						close(wardDone)
						startWard()
						continue loop
					case <-done:
						return
					}
				}
			}
		}()
		return heartbeat
	}
}

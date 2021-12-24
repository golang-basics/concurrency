package main

import "runtime"

func main() {
	// this will cause the main goroutine to be terminated, resulting in a deadlock
	// basically the scheduler will check for running go routines, if none are running it will result in
	// waiting forever on nothing to run => deadlock
	runtime.Goexit()
}

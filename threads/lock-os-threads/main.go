package main

import "runtime"

func main() {
	// yo use it with a go routine
	runtime.LockOSThread()
}

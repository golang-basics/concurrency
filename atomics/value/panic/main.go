package main

import (
	"sync/atomic"
)

func main() {
	var v atomic.Value
	// this will panic, saying we can't pass nil
	// v.Store(nil)
	v.Store(1)
	// this will panic, saying types must be consistent
	// v.Store("1")
}

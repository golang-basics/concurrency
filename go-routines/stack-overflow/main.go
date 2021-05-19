package main

import (
	"time"
)

func f() {
	f()
}

func main() {
	go f()
	time.Sleep(10 * time.Second)
}

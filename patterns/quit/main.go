package main

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, os.Interrupt)

	err := worker(quit)

	select {
	case e := <-err:
		fmt.Println("error happened in worker:", e)
	default:
		return
	}
}

func worker(quit chan os.Signal) chan error {
	ticker := time.NewTicker(500 * time.Millisecond)
	timeout := time.NewTimer(3 * time.Second)
	for {
		select {
		case <-timeout.C:
			err := make(chan error, 1)
			err <- errors.New("something wrong happened")
			return err
		case <-ticker.C:
			fmt.Println("doing some work")
		case <-quit:
			cleanup()
			fmt.Println("exiting")
			return nil
		}
	}
}

func cleanup() {
	fmt.Println("doing some cleanup before exiting")
}

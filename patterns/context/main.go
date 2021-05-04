package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func main() {
	// Couple of Gotchas about contexts here:
	// 1. Pass the same instance of context to all go routines.
	// 2. Wrapping a context with cancel in a context with timeout is quite error prone
	// So avoid having multiple cancellation methods.
	// 3. Try to create new context instance from the same initial context
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		err := operation3(ctx, 1)
		if err != nil {
			cancel()
		}
	}()
	operation1(ctx)
	operation2(ctx)

	ctx, cancel = context.WithTimeout(ctx, time.Second)
	longRunningOperation(ctx)
}

func operation1(ctx context.Context) {
	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Println("operation 1 - done")
	case <-ctx.Done():
		fmt.Println("operation 1 - halted")
	}
}

func operation2(ctx context.Context) {
	select {
	case <-time.After(300 * time.Millisecond):
		fmt.Println("operation 2 - done")
	case <-ctx.Done():
		fmt.Println("operation 2 - halted")
	}
}

func operation3(ctx context.Context, n int) error {
	time.Sleep(200 * time.Millisecond)
	if n < 0 {
		return errors.New("something bad happened")
	}
	return nil
}

func longRunningOperation(ctx context.Context) {
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("finally done")
	case <-ctx.Done():
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("long running operation - timed out")
		} else {
			fmt.Println("long running operation - halted")
		}
	}
}

// add value to context
// extract value from context

package main

import (
	"context"
	"fmt"
	"time"
)

const timeout = 3 * time.Second

func main() {
	now := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	request(ctx)
	fmt.Println("elapsed:", time.Since(now))
}

func request(ctx context.Context) {
	if deadline(ctx, 500*time.Millisecond) {
		return
	}

	select {
	case <-ctx.Done():
		return
	case <-time.After(500 * time.Millisecond):
		fmt.Println("calling controller")
		controller(ctx)
	}
}

func controller(ctx context.Context) {
	if deadline(ctx, 800*time.Millisecond) {
		return
	}

	select {
	case <-ctx.Done():
		return
	case <-time.After(800 * time.Millisecond):
		fmt.Println("calling service")
		service(ctx)
	}
}

func service(ctx context.Context) {
	if deadline(ctx, 900*time.Millisecond) {
		return
	}

	select {
	case <-ctx.Done():
		return
	case <-time.After(900 * time.Millisecond):
		fmt.Println("calling repository")
		repository(ctx)
	}
}

func repository(ctx context.Context) {
	if deadline(ctx, 900*time.Millisecond) {
		return
	}

	select {
	case <-ctx.Done():
		return
	case <-time.After(900 * time.Millisecond):
		fmt.Println("calling db")
		db(ctx)
	}
}

func db(ctx context.Context) {
	select {
	case <-ctx.Done():
		return
	default:
		fmt.Println("saved record in db")
	}
}

// deadline calculates ahead of time of context is about to timeout
// and returns early shaving a little bit of the execution time.
// If a concurrent process is guarded by a timeout,
// sometimes knowing the timeout beforehand, we can return early
// instead of waiting for the context to be cancelled.
func deadline(ctx context.Context, timeout time.Duration) bool {
	if deadline, ok := ctx.Deadline(); ok {
		if deadline.Sub(time.Now().Add(timeout)) <= 0 {
			return true
		}
	}
	return false
}

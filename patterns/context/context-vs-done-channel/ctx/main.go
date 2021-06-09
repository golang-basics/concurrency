package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		if err := welcome(ctx); err != nil {
			fmt.Println("could not print welcome:", err)
			cancel()
		}
	}()
	go func() {
		defer wg.Done()
		if err := goodbye(ctx); err != nil {
			fmt.Println("could not print goodbye:", err)
		}
	}()
	wg.Wait()
}

func welcome(ctx context.Context) error {
	msg, err := genWelcome(ctx)
	if err != nil {
		return err
	}
	fmt.Println(msg)
	return nil
}

func genWelcome(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	switch l, err := locale(ctx); {
	case err != nil:
		return "", err
	case l == "EN/US":
		return "welcome", nil
	}
	return "", errors.New("unsupported locale")
}

func goodbye(ctx context.Context) error {
	msg, err := genGoodbye(ctx)
	if err != nil {
		return err
	}
	fmt.Println(msg)
	return nil
}

func genGoodbye(ctx context.Context) (string, error) {
	switch l, err := locale(ctx); {
	case err != nil:
		return "", err
	case l == "EN/US":
		return "goodbye", nil
	}
	return "", errors.New("unsupported locale")
}

func locale(ctx context.Context) (string, error) {
	timeout := 5 * time.Second
	// return early if context work should be cancelled
	// instead of waiting for the context deadline to pass
	// As you can see the official docs say the same
	// https://github.com/golang/go/blob/master/src/context/context.go#L66
	if deadline, ok := ctx.Deadline(); ok {
		if deadline.Sub(time.Now().Add(timeout)) <= 0 {
			return "", context.DeadlineExceeded
		}
	}

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-time.After(timeout):
	}
	return "EN/US", nil
}

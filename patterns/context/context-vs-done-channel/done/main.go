package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

func main() {
	done := make(chan struct{})
	defer close(done)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		if err := welcome(done); err != nil {
			fmt.Println("could not print welcome:", err)
			return
		}
	}()
	go func() {
		defer wg.Done()
		if err := goodbye(done); err != nil {
			fmt.Println("could not print goodbye:", err)
			return
		}
	}()
	wg.Wait()
}

func welcome(done <-chan struct{}) error {
	msg, err := genWelcome(done)
	if err != nil {
		return err
	}
	fmt.Println(msg)
	return nil
}

func genWelcome(done <-chan struct{}) (string, error) {
	switch l, err := locale(done); {
	case err != nil:
		return "", err
	case l == "EN/US":
		return "welcome", nil
	}
	return "", errors.New("unsupported locale")
}

func goodbye(done <-chan struct{}) error {
	msg, err := genGoodbye(done)
	if err != nil {
		return err
	}
	fmt.Println(msg)
	return nil
}

func genGoodbye(done <-chan struct{}) (string, error) {
	switch l, err := locale(done); {
	case err != nil:
		return "", err
	case l == "EN/US":
		return "goodbye", nil
	}
	return "", errors.New("unsupported locale")
}

func locale(done <-chan struct{}) (string, error) {
	select {
	case <-done:
		return "", fmt.Errorf("cancelled")
	case <-time.After(5 * time.Second):
	}
	return "EN/US", nil
}

package main

import (
	"errors"
	"fmt"
)

var errNotFound = errors.New("resource not found")

type customError struct {
	message string
	err     error
}

func (e customError) Error() string {
	return e.message
}

func (e customError) Unwrap() error {
	return e.err
}

func main() {
	err := request()
	if !errors.As(err, &customError{}) {
		fmt.Println("something unexpected happened")
		return
	}
	if errors.Is(err, errNotFound) {
		fmt.Println("could not find resource")
	}
}

func request() error {
	return customError{
		message: "request error",
		err:     controller(),
	}
}

func controller() error {
	return fmt.Errorf("%w: %v", service(), "controller error")
}

func service() error {
	return fmt.Errorf("%w: %v", repo(), "service error")
}

func repo() error {
	return fmt.Errorf("%w: %v", errNotFound, "repo error")
}

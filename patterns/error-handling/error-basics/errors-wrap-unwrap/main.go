package main

import (
	"errors"
	"fmt"
	"os"
)

// The functions: As(), Is(), Unwrap() inside the errors package
// are available to use as of Go 1.13, which was inspired from:
// https://github.com/pkg/errors
// For more info on new error features in Go 1.13, check out:
// https://blog.golang.org/go1.13-errors
func main() {
	err := request()
	if err != nil {
		for err != nil {
			fmt.Println("err:", err)
			err = errors.Unwrap(err)
		}
		os.Exit(1)
	}
	fmt.Println("success")
}

func request() error {
	return fmt.Errorf("request error: %w", controller())
}

func controller() error {
	return fmt.Errorf("controller error: %w", service())
}

func service() error {
	return fmt.Errorf("service error: %w", repo())
}

func repo() error {
	return errors.New("repo error")
}

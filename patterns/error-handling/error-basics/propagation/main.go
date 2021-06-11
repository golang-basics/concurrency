package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

var errNegativeNumber = errors.New("negative number")

type validationError struct {
	message string
	err     error
}

func (e validationError) Error() string {
	return e.message
}

type httpError struct {
	message string
	status  int
}

func (e httpError) Error() string {
	return e.message
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("no argument provided")
	}
	n := os.Args[1]
	err := request(n)
	switch e := err.(type) {
	case nil:
		fmt.Println("[200] - success")
	case httpError:
		fmt.Printf("[%d] - %s\n", e.status, e.message)
	default:
		fmt.Println("something unexpected happened")
	}
}

func request(number string) error {
	err := service(number)
	switch e := err.(type) {
	case nil:
		return nil
		// return errors.New("something unexpected")
	case validationError:
		return httpError{
			message: e.message,
			status:  400,
		}
	default:
		return httpError{
			message: "internal server error",
			status:  500,
		}
	}
}

func service(number string) error {
	return validate(number)
}

func validate(number string) error {
	n, err := strconv.Atoi(number)
	if err != nil {
		return validationError{
			message: "could not convert argument to int",
			err:     err,
		}
	}
	if n < 0 {
		return validationError{
			message: "negative number provided",
			err:     errNegativeNumber,
		}
	}
	return nil
}

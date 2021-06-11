package main

import (
	"fmt"
	"log"
)

func main() {
	err := request(-2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("successful request")
}

type httpError struct {
	message    string
	statusCode int
}

func (e httpError) Error() string {
	return e.message
}

func request(n int) error {
	if n < 0 {
		return httpError{
			message:    "negative number provided",
			statusCode: 400,
		}
	}
	return nil
}

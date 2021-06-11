package main

import (
	"errors"
	"fmt"
	"runtime/debug"
)

func main() {
	err := request()
	var trace string

	for err != nil {
		fmt.Println("err:", err)
		e, ok := err.(customError)
		if ok {
			trace = e.stacktrace
		}
		err = errors.Unwrap(err)
	}

	fmt.Printf("stack trace:\n %s\n", trace)
}

type customError struct {
	err        error
	message    string
	stacktrace string
	context    map[string]interface{}
}

func (e customError) Error() string {
	return e.message
}

func (e customError) Unwrap() error {
	return e.err
}

func wrap(err error, message string, ctx map[string]interface{}) customError {
	return customError{
		err:        err,
		message:    message,
		stacktrace: string(debug.Stack()),
		context:    ctx,
	}
}

func request() error {
	return wrap(controller(), "request error", nil)
}

func controller() error {
	return wrap(service(), "controller error", nil)
}

func service() error {
	return wrap(repo(), "service error", nil)
}

func repo() error {
	e := errors.New("something bad happened")
	return wrap(e, "repo error", nil)
}

package main

import (
	"errors"
)

var errInvalidAge = errors.New("person's age must be between 1-100")

type HTTPError struct {
	Message string `json:"message"`
	Err     error  `json:"-"`
}

func (e HTTPError) Error() string {
	return e.Message
}

func (e HTTPError) Unwrap() error {
	return e.Err
}

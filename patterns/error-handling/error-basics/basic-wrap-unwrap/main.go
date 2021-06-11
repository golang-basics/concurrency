package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	err := request()
	if err != nil {
		for err != nil {
			fmt.Println("err:", err)
			err = unwrap(err)
		}
		os.Exit(1)
	}
	fmt.Println("success")
}

func wrap(err error, msg string) error {
	return fmt.Errorf("%s: %v", msg, err)
}

func unwrap(err error) error {
	es := strings.Split(err.Error(), ":")
	if len(es) == 1 {
		return nil
	}
	return errors.New(strings.TrimSpace(strings.Join(es[1:], ":")))
}

func request() error {
	return wrap(controller(), "error inside request")
}

func controller() error {
	return wrap(service(), "error inside controller")
}

func service() error {
	return wrap(repo(), "error inside service")
}

func repo() error {
	return errors.New("error inside repo")
}

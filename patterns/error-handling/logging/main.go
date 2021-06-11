package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	loggerInit()

	if len(os.Args) < 2 {
		log.Fatal("missing id argument")
	}

	id := os.Args[1]
	err := request(id)
	if err != nil {
		response(err)
		return
	}

	res := struct {
		Message string `json:"message"`
	}{"success"}
	response(res)
}

func response(v interface{}) {
	err := json.NewEncoder(os.Stdout).Encode(v)
	if err != nil {
		log.Fatalf("could not encode json: %v", err)
	}
}

func request(age string) error {
	a, err := strconv.Atoi(age)
	if err != nil {
		logError("could not save person with age: %s in controller\n", age)
		return HTTPError{
			Message: fmt.Sprintf("error converting %s to number", age),
			Err:     err,
		}
	}

	err = controller(a)
	if errors.Is(err, errInvalidAge) {
		return HTTPError{
			Message: "person's age is not correct, age must be between 1-100",
			Err:     err,
		}
	}

	if err != nil {
		logError("could not save person with age: %s from controller\n", age)
		return HTTPError{
			Message: "request error",
			Err:     err,
		}
	}

	logDebug("successfully called controller & saved person")
	return nil
}

func controller(age int) error {
	err := service(age)
	if err != nil {
		log.Printf("could not save person with age: %d in service\n", age)
		return fmt.Errorf("%w: %v", err, "controller error")
	}

	logDebug("successfully called service & saved person")
	return nil
}

func service(age int) error {
	err := repo(age)
	if err != nil {
		logError("could not save person with age: %d in repo\n", age)
		return fmt.Errorf("%w: %v", err, "service error")
	}

	logDebug("successfully called repo & saved person")
	return nil
}

func repo(age int) error {
	if age < 1 || age > 100 {
		logError("could not save person with age: %d", age)
		return errInvalidAge
	}

	logInfo("successfully saved person inside repo")
	return nil
}

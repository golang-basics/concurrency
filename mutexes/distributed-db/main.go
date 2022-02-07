package main

import (
	"log"

	"distributed-db/app"
)

func main() {
	a, err := app.New()
	if err != nil {
		log.Fatalf("could not create the app: %v", err)
	}

	err = a.Start()
	if err != nil {
		log.Fatalf("could not start the app: %v", err)
	}
}

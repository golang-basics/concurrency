package main

import (
	"log"
	"os"
	"syscall"

	"github.com/steevehook/http/app"
	"github.com/steevehook/http/worker"
)

func main() {
	a, err := app.Init()
	if err != nil {
		log.Fatalf("could not initialize application: %v", err)
	}

	w, err := worker.Init()
	if err != nil {
		log.Fatalf("could not initialize worker: %v", err)
	}

	go func() {
		if err := a.Start(); err != nil {
			log.Fatalf("could not start application: %v", err)
		}
	}()
	go func() {
		if err := w.Start(); err != nil {
			log.Fatalf("could not start worker: %v", err)
		}
	}()

	app.ListenToSignals([]os.Signal{os.Interrupt, syscall.SIGTERM}, a, w)
}

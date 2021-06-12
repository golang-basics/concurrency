package main

import (
	"log"
	"os"
	"syscall"

	"github.com/steevehook/http/app"
	"github.com/steevehook/http/db"
	"github.com/steevehook/http/repositories"
	"github.com/steevehook/http/worker"
)

func main() {
	d, err := db.Init()
	if err != nil {
		log.Fatalf("could not initialize database: %v", err)
	}

	repo := repositories.NewBookings(d)
	err = repo.Init(70)
	if err != nil {
		log.Fatalf("could not initialize repository: %v", err)
	}

	a, err := app.Init(repo)
	if err != nil {
		log.Fatalf("could not initialize application: %v", err)
	}

	w, err := worker.Init(repo)
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

	app.ListenToSignals([]os.Signal{os.Interrupt, syscall.SIGTERM}, a, w, d)
}

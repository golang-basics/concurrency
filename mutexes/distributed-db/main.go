package main

import (
	"context"
	"distributed-db/app"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	a, err := app.New()
	if err != nil {
		log.Fatalf("could not create the app: %v", err)
	}

	go func() {
		err = a.Start(ctx)
		if err != nil {
			log.Fatalf("could not start the app: %v", err)
		}
	}()

	select {
	case <-signals:
		cancel()
		err = a.Stop(ctx)
		if err != nil {
			log.Fatalf("could not stop the app: %v", err)
		}
	}
}

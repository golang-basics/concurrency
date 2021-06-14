package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	defer fmt.Println("done")
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	apiConn := open()

	var wg sync.WaitGroup
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			err := apiConn.readFile(ctx)
			if err != nil {
				log.Printf("could not read file: %v\n", err)
			}
			log.Println("read file")
		}()
	}
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			err := apiConn.httpCall(ctx)
			if err != nil {
				log.Printf("could not make http call: %v\n", err)
			}
			log.Println("http call")
		}()
	}

	wg.Wait()
}

func open() *apiConnection {
	return &apiConnection{
		rateLimiter: rate.NewLimiter(rate.Every(200*time.Millisecond), 1),
	}
}

type apiConnection struct {
	rateLimiter *rate.Limiter
}

func (a *apiConnection) readFile(ctx context.Context) error {
	if err := a.rateLimiter.Wait(ctx); err != nil {
		return err
	}
	// do some work here
	return nil
}

func (a *apiConnection) httpCall(ctx context.Context) error {
	if err := a.rateLimiter.Wait(ctx); err != nil {
		return err
	}
	// do some work here
	return nil
}

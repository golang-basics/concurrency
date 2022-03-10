package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"githubc.com/steevehook/log-reader/logging"
)

func main() {
	quit := make(chan os.Signal, 1)
	directoryFlag := flag.String("d", ".", "the directory where all the logs are stored")
	queryFlag := flag.String("q", "", "the query string to look for in the log files")
	limitFlag := flag.Int("n", 100, "the maximum number of logs to display")

	flag.Parse()
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	cfg := logging.ReaderConfig{
		Directory: *directoryFlag,
		Query:     *queryFlag,
		Limit:     *limitFlag,
	}
	fmt.Println(cfg)
	logReader, err := logging.NewReader(cfg)
	if err != nil {
		log.Fatalf("could not create log reader: %v", err)
	}

	go func() {
		err := logReader.Read(ctx, os.Stdout)
		if err != nil {
			log.Fatalf("could not read logs: %v", err)
		}

		quit <- os.Interrupt
	}()

	<-quit
	cancel()
}

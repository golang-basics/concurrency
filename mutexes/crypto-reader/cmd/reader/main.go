package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"githubc.com/steevehook/crypto-reader/crypto"
)

func main() {
	quit := make(chan os.Signal, 1)
	addressFlag := flag.String("address", "", "the crypto address to look for")
	directoryFlag := flag.String("directory", "", "the path to the crypto transaction files")
	intervalFlag := flag.Duration("interval", time.Hour, "the interval of transactions to look for")
	limitFlag := flag.Int("limit", 100, "the maximum number of transactions to read")

	flag.Parse()

	if *directoryFlag == "" {
		log.Fatal("provide the directory ('directory' flag) of all transaction files")
	}
	if *addressFlag == "" {
		log.Fatal("provide a crypto address ('address' flag) to read transactions for")
	}

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	cfg := crypto.TransactionsReaderConfig{
		Address:   *addressFlag,
		Interval:  *intervalFlag,
		Directory: *directoryFlag,
		Limit:     *limitFlag,
	}
	reader, err := crypto.NewTransactionsReader(cfg)
	if err != nil {
		log.Fatalf("could not create crypto transactions reader: %v", err)
	}

	go func() {
		err := reader.Read(ctx, os.Stdout)
		if err != nil {
			log.Fatalf("could not read crypto transactions: %v", err)
		}

		quit <- os.Interrupt
	}()

	<-quit
	cancel()
}

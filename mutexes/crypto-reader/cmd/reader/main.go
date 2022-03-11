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
	addrFlag := flag.String("addr", "", "the crypto address to look for")
	intervalFlag := flag.Duration("interval", time.Hour, "the interval of transactions to look for")
	dirFlag := flag.String("dir", ".", "the directory of crypto transaction files")
	nFlag := flag.Int("n", 100, "the maximum number of transactions to read")

	if *addrFlag == "" {
		log.Fatal("provide a crypto address to read transactions for")
	}

	flag.Parse()
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	cfg := crypto.TransactionsReaderConfig{
		Address:   *addrFlag,
		Interval:  *intervalFlag,
		Directory: *dirFlag,
		Limit:     *nFlag,
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

package app

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	"distributed-db/clients"
	"distributed-db/controllers"
	"distributed-db/models"
	"distributed-db/repositories"
	"distributed-db/services"
	"distributed-db/workers"
)

func New() (*App, error) {
	nodesMap := models.NodesMap{}
	port := flag.Int("port", 8080, "the port of the running server")
	dataDir := flag.String("data", "", "the data directory of the running server")
	flag.Var(&nodesMap, "node", "the list of nodes to talk to")

	flag.Parse()

	addr := fmt.Sprintf("localhost:%d", *port)
	if *dataDir == "" {
		*dataDir = fmt.Sprintf(".data/%s", addr)
	}

	delete(nodesMap, addr)
	if len(nodesMap) < 1 {
		return nil, fmt.Errorf("need at least 1 node to talk to")
	}

	nodes := models.NewNodes(addr, nodesMap)
	tokens := models.NewTokens(nodes, 256)
	cacheRepo := repositories.NewCache(*dataDir)
	httpClient := clients.NewHTTP(addr)
	svc := services.NewCache(cacheRepo, httpClient, tokens)
	router := controllers.NewRouter(svc)
	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}
	gossipWorker := workers.NewGossip(svc)
	streamerWorker := workers.NewStreamer(svc)
	a := &App{
		Server:         srv,
		GossipWorker:   gossipWorker,
		StreamerWorker: streamerWorker,
		cacheRepo:      cacheRepo,
	}

	return a, nil
}

type snapshotter interface {
	Snapshot() error
}

type App struct {
	Server         *http.Server
	GossipWorker   workers.Gossip
	StreamerWorker workers.Streamer
	cacheRepo      snapshotter
}

func (a App) Start(ctx context.Context) error {
	go a.GossipWorker.Start(ctx)
	go a.StreamerWorker.Start(ctx)

	log.Println("server started on address", a.Server.Addr)
	err := a.Server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

func (a App) Stop(ctx context.Context) error {
	log.Println("shutting down the http server")
	err := a.Server.Shutdown(ctx)
	if err != nil && err != context.Canceled {
		return fmt.Errorf("could not stop the http server: %w", err)
	}

	log.Println("taking a snapshot of the database")
	err = a.cacheRepo.Snapshot()
	if err != nil {
		return fmt.Errorf("could not take database snapshot: %w", err)
	}
	return nil
}

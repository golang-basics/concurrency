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
	nodes := models.Nodes{Map: map[string]struct{}{}}
	port := flag.Int("port", 8080, "the port of the running server")
	flag.Var(&nodes, "node", "the list of nodes to talk to")

	flag.Parse()
	if len(nodes.Map) < 1 {
		return nil, fmt.Errorf("need at least 1 node to talk to")
	}

	addr := fmt.Sprintf("localhost:%d", *port)
	nodes.CurrentNode = addr
	tokens := models.NewTokens(nodes, 256)
	cacheRepo := repositories.NewCache()
	httpClient := clients.NewHTTP(addr)
	svc := services.NewCache(cacheRepo, httpClient, tokens)
	router := controllers.NewRouter(svc)
	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}
	w := workers.NewGossip(svc)
	a := &App{
		Server:  srv,
		Worker:  w,
	}

	return a, nil
}

type App struct {
	Server  *http.Server
	Worker  workers.Gossip
}

func (a App) Start(ctx context.Context) error {
	go a.Worker.Start(ctx)

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
	return nil
}

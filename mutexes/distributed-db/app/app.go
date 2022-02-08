package app

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"distributed-db/clients"
	"distributed-db/controllers"
	"distributed-db/models"
	"distributed-db/repositories"
	"distributed-db/services"
	"distributed-db/workers"
)

func New() (*App, error) {
	peers := models.Peers{}
	port := flag.Int("port", 8080, "the port of the running server")
	flag.Var(&peers, "peer", "the list of peer servers to gossip to")

	flag.Parse()
	if len(peers) < models.MinimumPeers {
		return nil, fmt.Errorf("need at least %d peer servers", models.MinimumPeers)
	}

	addr := ":" + strconv.Itoa(*port)
	cacheRepo := repositories.NewCache()
	httpClient := clients.NewHTTP("localhost" + addr)
	svc := services.NewCache(cacheRepo, httpClient, peers)
	router := controllers.NewRouter(svc)
	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}
	w := workers.NewGossip(svc)
	a := &App{
		Server: srv,
		Worker: w,
	}

	return a, nil
}

type App struct {
	Server *http.Server
	Worker workers.Gossip
}

func (a App) Start() error {
	go a.Worker.Start()

	log.Println("server started on address", a.Server.Addr)
	err := a.Server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

package app

import (
	"flag"
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
	var peers models.StringList
	port := flag.Int("port", 8080, "the port of the running server")
	flag.Var(&peers, "peer", "the list of peer servers to gossip to")

	flag.Parse()

	addr := ":" + strconv.Itoa(*port)
	peersRepo, err := repositories.NewPeers(peers)
	if err != nil {
		return nil, err
	}
	cacheRepo := repositories.NewCache()
	httpClient := clients.NewHTTP("localhost" + addr)
	svc := services.NewCache(cacheRepo, peersRepo, httpClient)
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

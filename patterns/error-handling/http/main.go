package main

import (
	"log"
	"net/http"
	"time"

	"github.com/boltdb/bolt"
	"go.uber.org/zap"

	"github.com/steevehook/http/controllers"
	"github.com/steevehook/http/logging"
	"github.com/steevehook/http/repositories"
	"github.com/steevehook/http/services"
)

func main() {
	err := logging.Init()
	if err != nil {
		log.Fatalf("could not initialize logging: %v", err)
	}

	logger := logging.Logger()
	db, err := bolt.Open(
		"notes.db",
		0600,
		&bolt.Options{
			Timeout: 1 * time.Second,
		},
	)
	defer func() {
		err := db.Close()
		if err != nil {
			logger.Error("could not close bold db file database", zap.Error(err))
		}
	}()
	if err != nil {
		logger.Error("could not open bold db file database", zap.Error(err))
	}

	repo := repositories.NewNotes(db)
	service := services.NewNotes(repo)
	router := controllers.NewRouter(service)
	port := ":8080"

	logger.Info("server is up and running on port " + port)
	err = http.ListenAndServe(port, router)
	if err != nil {
		log.Fatalf("could run http server: %v", err)
	}
}

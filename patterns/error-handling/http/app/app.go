package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"go.uber.org/zap"

	"github.com/steevehook/http/controllers"
	"github.com/steevehook/http/logging"
	"github.com/steevehook/http/repositories"
	"github.com/steevehook/http/services"
)

type App struct {
	Server *http.Server
}

func Init(repo repositories.BookingsRepository) (*App, error) {
	err := logging.Init()
	if err != nil {
		return nil, fmt.Errorf("%v: %w", "could not initialize logger", err)
	}

	service := services.NewBookings(repo)
	router := controllers.NewRouter(service)
	app := &App{
		Server: &http.Server{
			Addr:         ":8080",
			Handler:      router,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
			ErrorLog:     logging.HTTPServerLogger(),
		},
	}
	return app, nil
}

func (a App) Start() error {
	logger := logging.Logger()
	logger.Info("server is up and running on port " + a.Server.Addr)

	err := a.Server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		logger.Error("could not listen and serve", zap.Error(err))
		return err
	}

	return nil
}

func (a App) Stop() error {
	logger := logging.Logger()
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	logger.Info("shutting down the http server")
	if err := a.Server.Shutdown(ctx); err != nil {
		logger.Error("error on server shutdown", zap.Error(err))
		return err
	}
	logger.Info("http server was successfully shut down")

	return nil
}

type stopper interface {
	Stop() error
}

// ListenToSignals listens for any incoming termination signals and shuts down the application(s)
func ListenToSignals(signals []os.Signal, apps ...stopper) {
	logger := logging.Logger()
	s := make(chan os.Signal, 1)
	signal.Notify(s, signals...)

	<-s
	var err error
	for _, a := range apps {
		err = a.Stop()
		if err != nil {
			logger.Error("stopping resulted in error", zap.Error(err))
		}
	}
	if err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}

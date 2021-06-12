package worker

import (
	"context"
	"time"

	"github.com/steevehook/http/db"
	"github.com/steevehook/http/logging"
	"github.com/steevehook/http/repositories"
)

type bookingsRepo interface {
	DeleteExpiredBookings(ctx context.Context) error
}

type Worker struct {
	done chan struct{}
	repo bookingsRepo
}

func Init(db db.DB) (*Worker, error) {
	repo := repositories.NewBookings(db)
	worker := &Worker{
		repo: repo,
		done: make(chan struct{}),
	}
	return worker, nil
}

func (w Worker) Start() error {
	logger := logging.Logger()
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-w.done:
			return nil
		case <-ticker.C:
			logger.Debug("deleting expired bookings")
		}
	}
}

func (w Worker) Stop() error {
	// write some stats to a file
	logger := logging.Logger()
	logger.Info("shutting worker down")
	close(w.done)
	logger.Info("worker was successfully shut down")
	return nil
}

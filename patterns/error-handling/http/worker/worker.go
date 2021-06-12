package worker

import (
	"context"
	"time"

	"go.uber.org/zap"

	"github.com/steevehook/http/logging"
	"github.com/steevehook/http/repositories"
)

const WorkPeriod = time.Hour

type bookingsRepo interface {
	DeleteExpiredBookings(ctx context.Context) (int, error)
}

// Worker represents the background worker that cleans expired bookings
type Worker struct {
	done chan struct{}
	repo bookingsRepo
}

// Init initializes the background worker
func Init(repo repositories.BookingsRepository) (*Worker, error) {
	worker := &Worker{
		repo: repo,
		done: make(chan struct{}),
	}
	return worker, nil
}

// Start starts the background worker
func (w Worker) Start() error {
	logger := logging.Logger()
	ticker := time.NewTicker(WorkPeriod)
	failures, maxFailures := 0, 5
	for {
		select {
		case <-w.done:
			return nil
		case <-ticker.C:
			logger.Info("deleting expired bookings")
			n, err := w.repo.DeleteExpiredBookings(context.Background())
			if err != nil {
				failures++
				logger.Debug("could not delete expired bookings", zap.Error(err))
				if failures == maxFailures {
					return err
				}
				continue
			}
			logger.Info("successfully deleted expired bookings", zap.Int("bookings", n))
		}
	}
}

// Stop stops the background worker
func (w Worker) Stop() error {
	logger := logging.Logger()
	logger.Info("shutting worker down")
	close(w.done)
	logger.Info("worker was successfully shut down")
	return nil
}

package db

import (
	"time"

	"github.com/boltdb/bolt"
	"go.uber.org/zap"

	"github.com/steevehook/http/logging"
)

func Init() (DB, error) {
	logger := logging.Logger()
	db, err := bolt.Open(
		"bookings.db",
		0600,
		&bolt.Options{
			Timeout: 1 * time.Second,
		},
	)
	if err != nil {
		logger.Error("could not open bolt database", zap.Error(err))
		return DB{}, err
	}

	return DB{db}, nil
}

type DB struct {
	*bolt.DB
}

func (d DB) Stop() error {
	logger := logging.Logger()
	logger.Info("closing the database")

	err := d.Close()
	if err != nil {
		logger.Error("could not close the database", zap.Error(err))
		return err
	}

	logger.Info("database was successfully closed")
	return nil
}

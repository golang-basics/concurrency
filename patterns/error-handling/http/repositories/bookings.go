package repositories

import (
	"context"
	"encoding/json"

	"github.com/boltdb/bolt"
	"go.uber.org/zap"

	"github.com/steevehook/http/logging"
	"github.com/steevehook/http/models"
)

const bucketName = "bookings"

type db interface {
	View(func(tx *bolt.Tx) error) error
	Update(func(tx *bolt.Tx) error) error
}

// NewBookings creates a new instance of BookingsRepository
func NewBookings(db db) BookingsRepository {
	return BookingsRepository{
		db: db,
	}
}

// BookingsRepository represents the Bookings repository that will interact with the database
type BookingsRepository struct {
	db db
}

// CreateBooking creates and saves a booking inside the database
func (r BookingsRepository) CreateBooking(ctx context.Context, booking models.Booking) error {
	logger := logging.Logger()
	return r.db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			logger.Error("could not create bucket", zap.Error(err))
			return err
		}

		bs, err := json.Marshal(booking)
		if err != nil {
			logger.Error("could not marshal booking", zap.Error(err))
			return err
		}

		return bucket.Put([]byte(booking.ID), bs)
	})
}

// GetBooking fetches a booking from the database
func (r BookingsRepository) GetBooking(ctx context.Context, id string) (models.Booking, error) {
	logger := logging.Logger()
	booking := models.Booking{}
	err := r.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		notFoundErr := models.ResourceNotFoundError{
			Message: "could not find booking with id: " + id,
		}
		if bucket == nil {
			return notFoundErr
		}

		bs := bucket.Get([]byte(id))
		if len(bs) == 0 {
			return notFoundErr
		}

		err := json.Unmarshal(bs, &booking)
		if err != nil {
			logger.Error("could not unmarshal booking", zap.Error(err))
			return err
		}

		return nil
	})
	if err != nil {
		logger.Error("could not fetch booking", zap.Error(err))
		return models.Booking{}, err
	}

	return booking, nil
}

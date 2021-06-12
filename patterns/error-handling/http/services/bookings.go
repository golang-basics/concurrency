package services

import (
	"context"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/steevehook/http/logging"
	"github.com/steevehook/http/models"
)

type repo interface {
	GetBooking(ctx context.Context, id string) (models.Booking, error)
	CreateBooking(ctx context.Context, booking models.Booking) error
}

// NewBookings creates a new instance of BookingService
func NewBookings(r repo) BookingService {
	return BookingService{
		repo: r,
	}
}

// BookingService represents the booking service that interacts with booking repositories
type BookingService struct {
	repo
}

// CreateBooking creates a booking from the repository
func (r BookingService) CreateBooking(ctx context.Context, req models.CreateBookingRequest) (models.Booking, error) {
	logger := logging.Logger()
	id := uuid.New()
	booking := models.Booking{
		ID:        id.String(),
		StartsAt:  req.Start,
		EndsAt:    req.End,
		CreatedAt: time.Now().UTC(),
	}

	err := r.repo.CreateBooking(ctx, booking)
	if err != nil {
		logger.Error("could not create booking", zap.Error(err))
		return models.Booking{}, err
	}

	return booking, nil
}

// GetBooking fetches a booking from the repository
func (r BookingService) GetBooking(ctx context.Context, req models.GetBookingRequest) (models.Booking, error) {
	logger := logging.Logger()

	booking, err := r.repo.GetBooking(ctx, req.ID)
	if err != nil {
		logger.Error("could not fetch booking", zap.Error(err))
		return models.Booking{}, err
	}

	return booking, nil
}

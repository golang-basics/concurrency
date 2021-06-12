package services

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/steevehook/http/logging"
	"github.com/steevehook/http/models"
)

type repo interface {
	GetBooking(ctx context.Context, id string) (models.Booking, error)
	CreateBooking(ctx context.Context, booking models.Booking) (models.Booking, error)
}

// NewBookings creates a new instance of BookingService
func NewBookings(r repo) BookingService {
	return BookingService{
		repo: r,
	}
}

// BookingService represents the booking service that interacts with booking repositories
type BookingService struct {
	repo repo
}

// GetBooking fetches a booking from the repository
func (s BookingService) GetBooking(ctx context.Context, req models.GetBookingRequest) (models.Booking, error) {
	logger := logging.Logger()
	_, err := uuid.Parse(req.ID)
	if err != nil {
		e := models.FormatValidationError{
			Message: fmt.Sprintf("invalid uuid: %s", req.ID),
		}
		return models.Booking{}, e
	}

	booking, err := s.repo.GetBooking(ctx, req.ID)
	if err != nil {
		logger.Error("could not fetch booking", zap.Error(err))
		return models.Booking{}, err
	}

	return booking, nil
}

// CreateBooking creates a booking from the repository
func (s BookingService) CreateBooking(ctx context.Context, req models.CreateBookingRequest) (models.Booking, error) {
	logger := logging.Logger()
	id := uuid.New()
	booking := models.Booking{
		ID:        id.String(),
		HotelID:   req.HotelID,
		StartsAt:  req.Start.UTC(),
		EndsAt:    req.End.UTC(),
		CreatedAt: time.Now().UTC(),
	}

	if time.Now().UTC().After(booking.StartsAt) {
		err := models.DataValidationError{
			Message: "start cannot be smaller than current time",
		}
		return models.Booking{}, err
	}
	if booking.StartsAt.Add(24 * time.Hour).After(booking.EndsAt) {
		err := models.DataValidationError{
			Message: "end must be at least 24 hours greater than start",
		}
		return models.Booking{}, err
	}

	booking, err := s.repo.CreateBooking(ctx, booking)
	if err != nil {
		logger.Error("could not create booking", zap.Error(err))
		return models.Booking{}, err
	}

	return booking, nil
}

package models

import (
	"time"
)

// CreateBookingRequest represents the request for creating a booking
type CreateBookingRequest struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

// GetBookingRequest represents the request for fetching a booking
type GetBookingRequest struct {
	ID string `json:"id"`
}

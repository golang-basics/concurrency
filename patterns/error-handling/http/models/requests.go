package models

import (
	"time"
)

// CreateBookingRequest represents the request for creating a booking
type CreateBookingRequest struct {
	Start   time.Time `json:"start"`    // i.e 2021-06-13T09:30:00Z
	End     time.Time `json:"end"`      // i.e 2021-06-14T09:30:00Z
	HotelID string    `json:"hotel_id"` // i.e default_hotel_id
}

// GetBookingRequest represents the request for fetching a booking
type GetBookingRequest struct {
	ID string `json:"id"`
}

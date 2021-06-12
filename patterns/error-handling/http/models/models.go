package models

import (
	"time"
)

const DefaultHotelID = "default_hotel_id"

// Booking represents the Booking model
type Booking struct {
	ID         string    `json:"id"`
	HotelID    string    `json:"hotel_id"`
	RoomNumber int       `json:"room_number"`
	CreatedAt  time.Time `json:"created_at"`
	StartsAt   time.Time `json:"starts_at"`
	EndsAt     time.Time `json:"ends_at"`
}

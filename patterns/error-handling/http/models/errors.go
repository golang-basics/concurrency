package models

import (
	"fmt"
)

// HTTPError represents a generic HTTP error
type HTTPError struct {
	Code    int    `json:"-"`
	Type    string `json:"type"`
	Message string `json:"message,omitempty"`
}

func (e HTTPError) Error() string {
	return fmt.Sprintf("http error: %s", e.Message)
}

// FormatValidationError is returned when request has invalid format
type FormatValidationError struct {
	Message string
}

func (e FormatValidationError) Error() string {
	return e.Message
}

// DataValidationError represents an error type for invalid data provision
type DataValidationError struct {
	Message string
}

func (e DataValidationError) Error() string {
	return e.Message
}

// InvalidJSONError is returned when request body can't be decoded from JSON
type InvalidJSONError struct {
	Message string
}

func (e InvalidJSONError) Error() string {
	return e.Message
}

// ResourceNotFoundError represents an error type for not found resources on the server
type ResourceNotFoundError struct {
	Message string
}

func (e ResourceNotFoundError) Error() string {
	if e.Message == "" {
		return "resource not found"
	}
	return e.Message
}

// HotelFullError is returned when no more bookings can be made
type HotelFullError struct {
	HotelID string
}

func (e HotelFullError) Error() string {
	return fmt.Sprintf("the hotel with id '%s' is full", e.HotelID)
}

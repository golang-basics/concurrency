package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	"go.uber.org/zap"

	"github.com/steevehook/http/logging"
	"github.com/steevehook/http/models"
	"github.com/steevehook/http/transport"
)

type bookingCreator interface {
	CreateBooking(ctx context.Context, req models.CreateBookingRequest) (models.Booking, error)
}

func createBooking(service bookingCreator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := logging.Logger()

		var req models.CreateBookingRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			logger.Error("could not decode json", zap.Error(err))
			e := models.InvalidJSONError{
				Message: "could not decode json",
			}
			transport.SendHTTPError(w, e)
			return
		}

		booking, err := service.CreateBooking(r.Context(), req)
		if err != nil {
			transport.SendHTTPError(w, err)
			return
		}

		transport.SendJSON(w, http.StatusCreated, booking)
	}
}

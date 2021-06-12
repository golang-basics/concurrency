package controllers

import (
	"context"
	"net/http"

	"github.com/steevehook/http/logging"
	"github.com/steevehook/http/models"
	"github.com/steevehook/http/transport"
)

type bookingGetter interface {
	GetBooking(ctx context.Context, req models.GetBookingRequest) (models.Booking, error)
}

func getBooking(service bookingGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := logging.Logger()
		id := routeParam(r, idRouteParam)
		req := models.GetBookingRequest{
			ID: id,
		}

		booking, err := service.GetBooking(r.Context(), req)
		if err != nil {
			transport.SendHTTPError(w, err)
			return
		}

		logger.Info("successfully fetched booking")
		transport.SendJSON(w, http.StatusOK, booking)
	}
}

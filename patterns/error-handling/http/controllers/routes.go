package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type service interface {
	bookingGetter
	bookingCreator
}

func NewRouter(svc service) *httprouter.Router {
	router := httprouter.New()

	router.Handler(http.MethodGet, "/bookings/:id", getBooking(svc))
	router.Handler(http.MethodPost, "/bookings", createBooking(svc))
	router.NotFound = notFound()

	return router
}

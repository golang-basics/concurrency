package controllers

import (
	"net/http"

	"github.com/steevehook/http/models"
	"github.com/steevehook/http/transport"
)

func notFound() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		transport.SendHTTPError(w, models.ResourceNotFoundError{})
	}
}

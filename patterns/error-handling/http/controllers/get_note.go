package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"

	"github.com/steevehook/http/models"
	"github.com/steevehook/http/transport"
)

type notesGetter interface {
	GetNote(ctx context.Context, req models.GetNoteRequest) (models.Note, error)
}

func getNote(service notesGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := routeParam(r, "id")
		_, err := uuid.Parse(id)
		if err != nil {
			e := models.FormatValidationError{
				Message: fmt.Sprintf("invalid uuid: %s", id),
			}
			transport.SendHTTPError(w, e)
			return
		}

		req := models.GetNoteRequest{
			ID: id,
		}
		note, err := service.GetNote(r.Context(), req)
		if err != nil {
			transport.SendHTTPError(w, err)
			return
		}

		transport.SendJSON(w, http.StatusOK, note)
	}
}

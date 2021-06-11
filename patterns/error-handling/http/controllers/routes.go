package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type service interface {
	notesGetter
	notesCreator
}

func NewRouter(svc service) *httprouter.Router {
	router := httprouter.New()

	router.Handler(http.MethodGet, "/notes/:id", getNote(svc))
	router.Handler(http.MethodPost, "/notes", createNote(svc))
	router.NotFound = notFound()

	return router
}

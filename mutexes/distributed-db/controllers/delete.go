package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"distributed-db/models"
)

type cacheRemover interface {
	Delete(keys []string)
}

func remove(svc cacheRemover) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.DeleteRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			log.Printf("could not decode delete request: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		svc.Delete(req.Keys)
		w.WriteHeader(http.StatusOK)
	}
}

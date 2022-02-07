package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"distributed-db/models"
)

type cacheGetter interface {
	Get(keys []string) []models.CacheItem
}

func get(svc cacheGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.GetRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			log.Printf("could not decode get request: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		values := svc.Get(req.Keys)

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(values)
		if err != nil {
			log.Printf("could not encode json: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

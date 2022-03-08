package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"distributed-db/models"
)

type cacheBatchSetter interface {
	SetBatch(items map[int]models.CacheItem) []models.CacheItem
}

func setBatch(svc cacheBatchSetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.SetBatchRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			log.Printf("could not decode set batch request: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		items := svc.SetBatch(req.Items)
		log.Printf("successfully stored batch")

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(items)
		if err != nil {
			log.Printf("could not encode set batch response: %v", err)
		}
	}
}

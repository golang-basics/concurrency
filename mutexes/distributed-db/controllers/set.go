package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"distributed-db/models"
)

type cacheSetter interface {
	Set(key, value string) (models.CacheItem, error)
}

func set(svc cacheSetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.SetRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			log.Printf("could not decode set request: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		item, err := svc.Set(req.Key, req.Value)
		if err != nil {
			log.Printf("could not store cache item: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		log.Printf("successfully stored record with key: %s on: %s", item.Key, item.Node)
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(item)
		if err != nil {
			log.Printf("could not encode set response: %v", err)
		}
	}
}

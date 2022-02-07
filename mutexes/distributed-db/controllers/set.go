package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

type cacheSetter interface {
	Set(key, value string)
}

type setRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func set(svc cacheSetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req setRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			log.Printf("could not decode set request: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		svc.Set(req.Key, req.Value)

		log.Printf("successfully stored record with key: %s", req.Key)
		w.WriteHeader(http.StatusOK)
	}
}

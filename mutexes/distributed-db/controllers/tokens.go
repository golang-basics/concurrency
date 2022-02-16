package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"distributed-db/models"
)

type tokensGetter interface {
	GetTokens() map[int]string
}

func tokens(svc tokensGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := models.TokensResponse{
			Tokens: svc.GetTokens(),
		}

		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("could not encode gossip response: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

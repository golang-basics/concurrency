package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"distributed-db/models"
)

type cacheSummaryResolver interface {
	ResolveSummary(peer string, summary models.Summary)
}

func gossip(svc cacheSummaryResolver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.GossipRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			log.Printf("could not decode gossip request: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		svc.ResolveSummary(r.Host, req.Summary)
	}
}

package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"distributed-db/models"
)

type tokensUpdater interface {
	UpdateTokens(node string, newNodes models.NodesMap, tokensChecksum string) (oldNodes models.NodesMap, err error)
}

func gossip(svc tokensUpdater) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.GossipRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			log.Printf("could not decode gossip request: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		oldNodes, err := svc.UpdateTokens(r.Host, req.Nodes, req.TokensChecksum)
		if err != nil {
			log.Printf("could not update tokens: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		res := models.GossipResponse{Nodes: oldNodes}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("could not encode gossip response: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

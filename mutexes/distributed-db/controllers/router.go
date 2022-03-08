package controllers

import (
	"net/http"
)

type CacheService interface {
	cacheGetter
	cacheSetter
	cacheBatchSetter
	cacheRemover
	tokensGetter
	tokensUpdater
}

func NewRouter(svc CacheService) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/get", get(svc))
	mux.HandleFunc("/set", set(svc))
	mux.HandleFunc("/delete", remove(svc))
	// add a middleware that does not allow
	// requests that do not come from other nodes
	mux.HandleFunc("/set/batch", setBatch(svc))
	mux.HandleFunc("/gossip", gossip(svc))
	mux.HandleFunc("/tokens", tokens(svc))

	return mux
}

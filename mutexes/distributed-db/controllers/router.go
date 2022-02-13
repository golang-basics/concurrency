package controllers

import (
	"net/http"
)

type CacheService interface {
	cacheGetter
	cacheSetter
	tokensGetter
	tokensUpdater
}

func NewRouter(svc CacheService) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/set", set(svc))
	mux.HandleFunc("/get", get(svc))
	mux.HandleFunc("/gossip", gossip(svc))
	mux.HandleFunc("/tokens", tokens(svc))

	return mux
}

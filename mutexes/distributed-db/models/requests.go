package models

type GetRequest struct {
	Keys []string `json:"keys"`
}

type GossipRequest struct {
	Summary Summary `json:"summary"`
}

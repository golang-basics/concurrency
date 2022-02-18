package models

type GossipResponse struct {
	Nodes map[string]int `json:"nodes"`
}

type TokensResponse struct {
	Tokens TokenMappings `json:"tokens"`
}

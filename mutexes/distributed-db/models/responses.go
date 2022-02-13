package models

type GossipResponse struct {
	Nodes []string `json:"nodes"`
}

type TokensResponse struct {
	Tokens TokenMappings `json:"tokens"`
}

package models

type GetRequest struct {
	Keys []string `json:"keys"`
	// how many reads before returning
	ConsistencyLevel int `json:"-"`
}

type SetRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	// how many copies for this cache item
	ReplicationFactor int `json:"-"`
	// how many writes before returning
	ConsistencyLevel int `json:"-"`
}

type GossipRequest struct {
	Nodes          []string `json:"nodes"`
	TokensChecksum string   `json:"tokens_checksum"`
}

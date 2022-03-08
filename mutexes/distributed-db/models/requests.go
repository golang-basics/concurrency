package models

import (
	"time"
)

type GetRequest struct {
	Keys []string `json:"keys"`
	// how many reads before returning (replication factor > 1) => TO BE IMPLEMENTED
	ConsistencyLevel int `json:"-"`
}

type DeleteRequest struct {
	Keys []string `json:"keys"`
	// how many deletes before returning (replication factor > 1) => TO BE IMPLEMENTED
	ConsistencyLevel int `json:"-"`
}

type SetRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	// how many copies for this cache item [TO BE IMPLEMENTED]
	ReplicationFactor int `json:"-"`
	// how many writes before returning (replication factor > 1) => TO BE IMPLEMENTED
	ConsistencyLevel int `json:"-"`
	// for short-lived records [TO BE IMPLEMENTED]
	TTL time.Duration `json:"-"`
}

type SetBatchRequest struct {
	Items map[int]CacheItem
	// how many copies for this cache item [TO BE IMPLEMENTED]
	ReplicationFactor int `json:"-"`
	// how many writes before returning (replication factor > 1) => TO BE IMPLEMENTED
	ConsistencyLevel int `json:"-"`
}

type GossipRequest struct {
	Nodes          map[string]int `json:"nodes"`
	TokensChecksum string         `json:"tokens_checksum"`
}

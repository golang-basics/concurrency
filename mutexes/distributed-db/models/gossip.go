package models

import (
	"time"
)

type GossipMessage struct {
	CreatedAt time.Time `json:"created_at"`
	Peers     Peers `json:"peers"`
}

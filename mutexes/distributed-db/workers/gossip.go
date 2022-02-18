package workers

import (
	"context"
	"log"
	"time"
)

const gossipPeriod = 3 * time.Second

type gossiper interface {
	Gossip()
}

func NewGossip(svc gossiper) Gossip {
	return Gossip{
		svc: svc,
	}
}

type Gossip struct {
	svc gossiper
}

func (g *Gossip) Start(ctx context.Context) {
	log.Println("gossip worker started successfully")

	for {
		select {
		case <-ctx.Done():
			log.Println("stopping the gossip worker")
			return
		case <-time.NewTicker(gossipPeriod).C:
			g.svc.Gossip()
		}
	}
}

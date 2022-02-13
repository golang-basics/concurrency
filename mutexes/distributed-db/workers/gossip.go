package workers

import (
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

func (g *Gossip) Start() {
	log.Println("worker started successfully")
	for {
		g.svc.Gossip()
		time.Sleep(gossipPeriod)
	}
}

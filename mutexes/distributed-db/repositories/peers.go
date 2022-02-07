package repositories

import (
	"errors"
)

const minimumPeers = 2

type Peers map[string]struct{}

func NewPeers(peersList []string) (*Peers, error) {
	if len(peersList) < 2 {
		return nil, errors.New("need at least 2 peer servers")
	}

	peers := Peers{}
	for _, p := range peersList {
		peers[p] = struct{}{}
	}

	return &peers, nil
}

func (p Peers) Add(peer string) {
	p[peer] = struct{}{}
}

func (p Peers) Random() []string {
	peers := make([]string, 0, minimumPeers)
	for peer := range p {
		peers = append(peers, peer)
		if len(peers) == minimumPeers {
			break
		}
	}

	return peers
}

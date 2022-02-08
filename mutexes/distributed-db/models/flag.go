package models

import (
	"strings"
)

const MinimumPeers = 2

// rename to Nodes
type Peers map[string]struct{}

func (p Peers) Set(s string) error {
	p.Add(s)
	return nil
}

func (p Peers) String() string {
	return strings.Join(p.List(len(p)), ",")
}

func (p Peers) Add(s string) {
	p[s] = struct{}{}
}

func (p Peers) List(n int) []string {
	i, keys := 0, make([]string, 0, len(p))
	for k := range p {
		if i == n {
			break
		}
		keys = append(keys, k)
		i++
	}
	return keys
}

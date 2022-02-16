package models

import (
	"strings"
)

type Nodes struct {
	Map         map[string]struct{}
	CurrentNode string
}

func (n Nodes) Set(s string) error {
	n.Add(s)
	return nil
}

func (n Nodes) String() string {
	return strings.Join(n.List(len(n.Map)), ",")
}

func (n Nodes) Add(nodes ...string) {
	for _, node := range nodes {
		if node == n.CurrentNode {
			continue
		}
		// update some kind of UpdateAt field for nodes health check
		// change from map[string]struct{} to map[string]int => states: Up/Down
		// a node is Down if it hasn't sent a gossip request in 10 seconds
		n.Map[node] = struct{}{}
	}
}

func (n Nodes) WithCurrentNode() map[string]struct{} {
	nodes := map[string]struct{}{n.CurrentNode: {}}
	for s := range n.Map {
		nodes[s] = struct{}{}
	}
	return nodes
}

func (n Nodes) List(a int) []string {
	i, keys := 0, make([]string, 0, len(n.Map))
	for k := range n.Map {
		if i == a {
			break
		}
		keys = append(keys, k)
		i++
	}
	return keys
}

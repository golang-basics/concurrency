package models

import (
	"strings"
)

type Nodes struct {
	Map       map[string]struct{}
	CurrentNode string
}

func (n Nodes) Set(s string) error {
	n.Add(s)
	return nil
}

func (n Nodes) String() string {
	return strings.Join(n.List(len(n.Map)), ",")
}

func (n Nodes) Add(node string) {
	if node == n.CurrentNode {
		return
	}
	n.Map[node] = struct{}{}
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

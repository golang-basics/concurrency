package models

import (
	"fmt"
	"strings"
)

type NodesMap map[string]int

func (m NodesMap) Set(element string) error {
	m[element] = NodeStatusUp
	return nil
}

func (m NodesMap) String() string {
	list := make([]string, 0, len(m))
	for node, status := range m {
		list = append(list, fmt.Sprintf("%s:%d", node, status))
	}
	return strings.Join(list, ",")
}

package models

import (
	"log"
	"time"
)

const (
	NodeStatusUp   = 1
	NodeStatusDown = 0
	gossipDeadline = 15 * time.Second
	retryPeriod    = time.Minute
	maxFails       = 5
)

func NewNodes(currentNode string, nodeMap NodesMap) *Nodes {
	nodes := &Nodes{
		current:      currentNode,
		nodesStatus:  NodesMap{},
		gossipStatus: map[string]time.Time{},
		nodeFails:    map[string]int{},
		nodeRetries:  map[string]time.Time{},
	}
	nodes.Set(nodeMap)
	go nodes.retry()
	return nodes
}

type Nodes struct {
	current      string
	nodesStatus  NodesMap
	gossipStatus map[string]time.Time
	nodeFails    map[string]int
	nodeRetries  map[string]time.Time
}

func (n Nodes) Current() string {
	return n.current
}

func (n Nodes) Fail(node string) {
	n.nodeFails[node]++
	if n.nodeFails[node] >= maxFails {
		log.Printf("node: %s is down, retryin in: %v", node, retryPeriod)
		n.nodesStatus[node] = NodeStatusDown
		n.nodeRetries[node] = time.Now().UTC()
	}
}

func (n Nodes) retry() {
	for {
		time.Sleep(time.Second)
		for node, lastTried := range n.nodeRetries {
			now := time.Now().UTC()
			if now.Sub(lastTried) >= retryPeriod {
				log.Printf("retrying gossip on node: %s", node)
				n.nodeFails[node] = 0
				n.nodesStatus[node] = NodeStatusUp
				delete(n.nodeRetries, node)
				n.gossipStatus[node] = time.Now().UTC()
			}
		}
	}
}

func (n Nodes) Gossip(node string) {
	n.gossipStatus[node] = time.Now().UTC()
}

func (n Nodes) Set(nodesStatus NodesMap) {
	for node, status := range nodesStatus {
		if node == n.current {
			continue
		}
		n.nodesStatus[node] = status
	}
}

func (n Nodes) Map() NodesMap {
	nodes := NodesMap{n.current: NodeStatusUp}
	for node, lastGossip := range n.gossipStatus {
		now := time.Now().UTC()
		if now.Sub(lastGossip) < gossipDeadline {
			nodes[node] = NodeStatusUp
			continue
		}
		nodes[node] = NodeStatusDown
	}
	return nodes
}

func (n Nodes) List(x int) []string {
	return n.list(x, false)
}

func (n Nodes) ListAll() []string {
	return n.list(len(n.nodesStatus), false)
}

func (n Nodes) ListActive(x int) []string {
	return n.list(x, true)
}

func (n Nodes) list(x int, filterByNodeStatusUp bool) []string {
	i, nodeList := 0, make([]string, 0, len(n.nodesStatus))
	for node, status := range n.nodesStatus {
		if i == x {
			break
		}

		if filterByNodeStatusUp && status != NodeStatusUp {
			continue
		}

		nodeList = append(nodeList, node)
		i++
	}
	return nodeList
}

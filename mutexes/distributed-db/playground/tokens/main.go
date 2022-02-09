package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"hash/fnv"
	"math"
	"math/rand"
	"sort"
	"time"
)

func main() {
	nodes := []Node{
		{
			ID: 1,
			Name: "server1",
			Address: "localhost:8081",
		},
		{
			ID: 2,
			Name: "server2",
			Address: "localhost:8082",
		},
		{
			ID: 3,
			Name: "server3",
			Address: "localhost:8083",
		},
		{
			ID: 4,
			Name: "server4",
			Address: "localhost:8084",
		},
		{
			ID: 5,
			Name: "server5",
			Address: "localhost:8085",
		},
	}
	tokens := NewTokens(nodes, 256)

	newNode := Node{
		ID: 6,
		Name: "server6",
		Address: "localhost:8086",
	}
	tokens.AddNode(newNode)

	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("k%d", i+1)
		sum := hash(key)
		fmt.Println(key, sum, tokens.GetNode(sum).Name)
	}

	// the message is quite big,
	// but it is exchanged only 1 time when a node joins
	bs, _ := json.Marshal(tokens.Mappings)
	// sending only the checksum can save a lot of space
	fmt.Printf("%x", md5.Sum(bs))
}

func NewTokens(nodes []Node, numberOfTokenRanges int) Tokens {
	numberOfNodes := len(nodes)
	tokenRange := uint64(math.MaxInt / numberOfNodes / numberOfTokenRanges)
	ranges := make([]uint64, 0, numberOfNodes*numberOfTokenRanges)
	nodesMap := map[int]Node{}
	for i := 0; i < numberOfNodes; i++ {
		nodesMap[nodes[i].ID] = nodes[i]
		for j := numberOfTokenRanges * i; j < numberOfTokenRanges*(i+1); j++ {
			r := tokenRange * uint64(j + 1)
			ranges = append(ranges, r)
		}
	}

	randomRanges := append([]uint64{}, ranges...)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(randomRanges), func(i, j int) {
		randomRanges[i], randomRanges[j] = randomRanges[j], randomRanges[i]
	})

	i, mappings := 0, map[uint64]int{}
	for _, r := range randomRanges {
		mappings[r] = nodes[i].ID
		i++
		if i == numberOfNodes {
			i = 0
		}
	}

	tokens := Tokens{
		Ranges:   ranges,
		Mappings: mappings,
		Nodes: nodesMap,
		NumberOfTokenRanges: numberOfTokenRanges,
	}
	return tokens
}

type Tokens struct {
	Ranges              []uint64
	Mappings            map[uint64]int
	Nodes               map[int]Node
	NumberOfTokenRanges int
	rangeMax uint64
}

func (t *Tokens) GetNode(token uint64) Node {
	idx := sort.Search(len(t.Ranges)-1, func(i int) bool {
		return t.Ranges[i] >= token
	})
	nodeID := t.Mappings[t.Ranges[idx]]
	return t.Nodes[nodeID]
}

func (t *Tokens) AddNode(node Node) {
	t.Nodes[node.ID] = node
	tokenRange := uint64(math.MaxInt / len(t.Nodes) / t.NumberOfTokenRanges)
	newRanges := make([]uint64, 0, t.NumberOfTokenRanges)
	for i := 0; i < len(t.Ranges)+t.NumberOfTokenRanges; i++ {
		if i < len(t.Ranges) {
			r := t.Ranges[i]
			decrement := r - tokenRange*(uint64(i+1))
			newRange := r - decrement
			srv := t.Mappings[r]

			delete(t.Mappings, r)
			t.Mappings[newRange] = srv
			t.Ranges[i] = newRange
		} else {
			newRange := tokenRange * uint64(i+1)
			t.Mappings[newRange] = len(t.Nodes)
			newRanges = append(newRanges, newRange)
		}
	}
	t.Ranges = append(t.Ranges, newRanges...)
}

type Node struct {
	ID   int
	Name string
	Address string
}

func hash(s string) uint64 {
	h := fnv.New64a()
	_, _ = h.Write([]byte(s))
	return h.Sum64()
}

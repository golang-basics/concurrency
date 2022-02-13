package models

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"hash/fnv"
	"log"
	"math/rand"
	"sort"
	"time"
)

type TokenMappings map[uint64]string

type Tokens struct {
	Mappings            TokenMappings
	Nodes               Nodes
	ranges              []uint64
	numberOfTokenRanges int
}

func NewTokens(nodes Nodes, numberOfTokenRanges int) Tokens {
	numberOfNodes := len(nodes.Map)
	tokenRange := uint64(200 / numberOfNodes / numberOfTokenRanges)
	ranges := make([]uint64, 0, numberOfNodes*numberOfTokenRanges)
	for i := 0; i < numberOfNodes; i++ {
		for j := numberOfTokenRanges * i; j < numberOfTokenRanges*(i+1); j++ {
			r := tokenRange * uint64(j+1)
			// produces a sorted ranges slice => needed for searching
			ranges = append(ranges, r)
		}
	}

	randomRanges := append([]uint64{}, ranges...)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(randomRanges), func(i, j int) {
		randomRanges[i], randomRanges[j] = randomRanges[j], randomRanges[i]
	})

	i, mappings := 0, map[uint64]string{}
	nodeList := nodes.List(len(nodes.Map))
	for _, r := range randomRanges {
		mappings[r] = nodeList[i]
		i++
		if i == numberOfNodes {
			i = 0
		}
	}

	tokens := Tokens{
		Mappings:            mappings,
		Nodes:               nodes,
		ranges:              ranges,
		numberOfTokenRanges: numberOfTokenRanges,
	}
	return tokens
}

func (t *Tokens) GetNode(token uint64) string {
	idx := sort.Search(len(t.ranges)-1, func(i int) bool {
		return t.ranges[i] >= token
	})
	node := t.Mappings[t.ranges[idx]]
	return node
}

func (t *Tokens) AddNode(node string) {
	_, ok := t.Nodes.Map[node]
	if ok || node == t.Nodes.CurrentNode {
		return
	}
	tokenRange := uint64(200 / len(t.Nodes.Map) / t.numberOfTokenRanges)
	newRanges := make([]uint64, 0, t.numberOfTokenRanges)
	for i := 0; i < len(t.ranges)+t.numberOfTokenRanges; i++ {
		if i < len(t.ranges) {
			r := t.ranges[i]
			decrement := r - tokenRange*(uint64(i+1))
			newRange := r - decrement
			srv := t.Mappings[r]

			delete(t.Mappings, r)
			t.Mappings[newRange] = srv
			t.ranges[i] = newRange
		} else {
			newRange := tokenRange * uint64(i+1)
			t.Mappings[newRange] = node
			newRanges = append(newRanges, newRange)
		}
	}
	t.ranges = append(t.ranges, newRanges...)
}

func (t *Tokens) Checksum() string {
	bs, err := json.Marshal(t.Mappings)
	if err != nil {
		log.Printf("could not marshal token mappings: %v", err)
		return ""
	}
	sum := md5.Sum(bs)
	return fmt.Sprintf("%x", sum)
}

func hash(s string) uint64 {
	h := fnv.New64a()
	_, _ = h.Write([]byte(s))
	return h.Sum64()
}

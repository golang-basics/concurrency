package models

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"hash/fnv"
	"log"
	"math"
	"math/rand"
	"reflect"
	"sort"
	"time"
)

func NewTokens(nodes *Nodes, numberOfTokenRanges int) *Tokens {
	nodeList := append([]string{nodes.Current()}, nodes.ListAll()...)
	tokenRange := math.MaxInt / len(nodeList) / numberOfTokenRanges
	ranges := make([]int, 0, len(nodeList)*numberOfTokenRanges)
	for i := 0; i < len(nodeList); i++ {
		for j := numberOfTokenRanges * i; j < numberOfTokenRanges*(i+1); j++ {
			r := tokenRange * (j + 1)
			// produces a sorted ranges slice => needed for searching
			ranges = append(ranges, r)
		}
	}

	randomRanges := append([]int{}, ranges...)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(randomRanges), func(i, j int) {
		randomRanges[i], randomRanges[j] = randomRanges[j], randomRanges[i]
	})

	i, mappings := 0, map[int]string{}
	for _, r := range randomRanges {
		mappings[r] = nodeList[i]
		i++
		if i == len(nodeList) {
			i = 0
		}
	}

	tokens := &Tokens{
		Mappings:            mappings,
		Nodes:               nodes,
		ranges:              ranges,
		numberOfTokenRanges: numberOfTokenRanges,
	}
	return tokens
}

type TokenMappings map[int]string

type Tokens struct {
	Mappings            TokenMappings
	ForeignTokens       TokenMappings
	Nodes               *Nodes
	ranges              []int
	numberOfTokenRanges int
}

func (t *Tokens) GetNode(token int) string {
	idx := sort.SearchInts(t.ranges, token)
	node := t.Mappings[t.ranges[idx]]
	return node
}

func (t *Tokens) SetForeignTokens(items map[int]CacheItem, node string) {
	for token, _ := range items {
		t.ForeignTokens[token] = node
	}
}

func (t *Tokens) DeleteForeignToken(token int) {
	delete(t.ForeignTokens, token)
}

func (t *Tokens) Merge(mappings map[int]string) {
	newMappings := map[int]string{}
	nodes := t.Nodes.Map()
	ranges := t.ranges
	m1, m2 := mappings, t.Mappings
	if len(mappings) < len(t.Mappings) {
		m1 = t.Mappings
		m2 = mappings
	}

	newNodes := map[string]int{}
	newRanges := make([]int, 0)
	for _, s := range mappings {
		newNodes[s] = NodeStatusUp
	}
	for r := range m1 {
		newRanges = append(newRanges, r)
	}
	sort.Ints(newRanges)
	ranges = newRanges
	nodes = newNodes

	if reflect.DeepEqual(nodes, t.Nodes.Map) {
		return
	}

	for s := range nodes {
		_, ok := t.Nodes.Map()[s]
		if ok {
			delete(nodes, s)
		}
	}

	i := 0
	numberOfNodes := len(nodes) + len(t.Nodes.Map())
	tokenRange := math.MaxInt / numberOfNodes / t.numberOfTokenRanges
	m1Nodes := map[string]struct{}{}
	for r, s := range m1 {
		m1Nodes[s] = struct{}{}
		factor := sort.SearchInts(ranges, r) + 1
		newMappings[factor*tokenRange] = s
		i++
	}
	for _, s := range m2 {
		_, ok := m1Nodes[s]
		if ok {
			continue
		}
		i++
		newMappings[(i)*tokenRange] = s
	}
	t.Mappings = newMappings
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

func HashKey(s string) uint64 {
	h := fnv.New64a()
	_, _ = h.Write([]byte(s))
	return h.Sum64()
}

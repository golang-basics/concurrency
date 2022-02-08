package main

import (
	"fmt"
	"hash/fnv"
	"math"
	"math/rand"
	"sort"
	"time"
)

func main() {
	tokens := NewTokens(5, 256)
	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("k%d", i+1)
		sum := hash(key)
		fmt.Println(key, sum, tokens.Get(sum))
	}
}

func NewTokens(numberOfNodes, numberOfTokenRanges int) Tokens {
	servers := make([]string, numberOfNodes)
	tokenRange := math.MaxInt / numberOfNodes / numberOfTokenRanges
	ranges := make([]uint64, 0, numberOfNodes*numberOfTokenRanges)
	for i := 0; i < numberOfNodes; i++ {
		servers[i] = fmt.Sprintf("server%d", i+1)
		for j := numberOfTokenRanges * i; j < numberOfTokenRanges*(i+1); j++ {
			r := tokenRange * (j + 1)
			ranges = append(ranges, uint64(r))
		}
	}

	randomRanges := append([]uint64{}, ranges...)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(randomRanges), func(i, j int) {
		randomRanges[i], randomRanges[j] = randomRanges[j], randomRanges[i]
	})

	i, mappings := 0, map[uint64]string{}
	for _, r := range randomRanges {
		mappings[r] = servers[i]
		i++
		if i == numberOfNodes {
			i = 0
		}
	}

	tokens := Tokens{
		Ranges:   ranges,
		Mappings: mappings,
	}
	return tokens
}

type Tokens struct {
	Ranges   []uint64
	Mappings map[uint64]string
}

func (t *Tokens) Get(n uint64) string {
	idx := sort.Search(len(t.Ranges)-1, func(i int) bool {
		return t.Ranges[i] >= n
	})
	return t.Mappings[t.Ranges[idx]]
}

func hash(s string) uint64 {
	h := fnv.New64a()
	_, _ = h.Write([]byte(s))
	return h.Sum64()
}

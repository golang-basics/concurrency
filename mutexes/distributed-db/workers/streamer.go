package workers

import (
	"context"
	"log"
	"time"

	"distributed-db/models"
)

const streamPeriod = 10 * time.Second

type streamer interface {
	Stream(retryStreams map[string]map[int]models.CacheItem) map[string]map[int]models.CacheItem
}

func NewStreamer(svc streamer) Streamer {
	return Streamer{
		svc:          svc,
		retryStreams: map[string]map[int]models.CacheItem{},
	}
}

type Streamer struct {
	svc          streamer
	retryStreams map[string]map[int]models.CacheItem
}

func (s *Streamer) Start(ctx context.Context) {
	itemsToRetry := 0
	log.Println("streamer worker started successfully")

	for {
		select {
		case <-ctx.Done():
			log.Println("stopping the streamer worker")
			return
		case <-time.NewTicker(streamPeriod).C:
			if itemsToRetry > 0 {
				log.Printf("retrying to stream %d item(s)", itemsToRetry)
			}

			failedBatches, failedItems := s.svc.Stream(s.retryStreams), 0
			for node, batchItems := range failedBatches {
				if s.retryStreams[node] == nil {
					s.retryStreams[node] = map[int]models.CacheItem{}
				}
				for token, item := range batchItems {
					s.retryStreams[node][token] = item
					failedItems++
				}
			}
			itemsToRetry = failedItems
		}
	}
}

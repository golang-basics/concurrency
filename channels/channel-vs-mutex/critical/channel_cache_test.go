package critical

type set struct {
	key   string
	value string
	done  chan struct{}
}

type channelCache struct {
	cache   map[string]string
	setChan chan set
}

func newChannelCache() *channelCache {
	cache := &channelCache{
		cache:   map[string]string{},
		setChan: make(chan set),
	}
	go func() {
		for {
			select {
			case op := <-cache.setChan:
				cache.cache[op.key] = op.value
				op.done <- struct{}{}
			}
		}
	}()
	return cache
}

func (c channelCache) set(key, value string) {
	op := set{
		key:   key,
		value: value,
		done:  make(chan struct{}),
	}
	c.setChan <- op
	<-op.done
}

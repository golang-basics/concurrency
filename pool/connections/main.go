package main

import "sync"

func main() {

}

// limit connections example
// investigate whether to Put the object back in the pool
// add concurrent example when Putting the object back in
func newCache() *cache {
	return &cache{
		pool: &sync.Pool{
			New: func() interface{} {
				return nil
			},
		},
	}
}

type cache struct {
	pool *sync.Pool
}

func (c *cache) set() {
	//c.pool.Put()
}

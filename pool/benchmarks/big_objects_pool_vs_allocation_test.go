package benchmarks

import (
	"strings"
	"sync"
	"testing"
)

type bigObject struct {
	bigString string
}

func newBigObject() interface{} {
	return bigObject{
		bigString: strings.Repeat("a", 1e9),
	}
}

func BenchmarkBigObjectAllocate(b *testing.B) {
	var bigObj bigObject
	// imagine we have to serve N clients
	// every time we serve a client we need a giant object to operate with
	// here we inefficiently create this object for every client
	// which allocates a ton of memory and also gives GC a hard time cleaning up
	// thus affecting the overall performance
	for i := 0; i < b.N; i++ {
		bigObj = newBigObject().(bigObject)
		bigObj = bigObj
	}
}

func BenchmarkBigObjectPool(b *testing.B) {
	pool := &sync.Pool{
		New: newBigObject,
	}
	// kinda works weird with go routines
	for i := 0; i < b.N; i++ {
		obj := pool.Get()
		pool.Put(obj)
	}
}

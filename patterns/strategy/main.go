package main

import (
	"github.com/goocarry/wb-internship/patterns/strategy/cache"
)

func main() {
	lfuAlgo := &cache.LfuAlgo{}
	c := cache.Cache{Storage: make(map[string]string), 
		EvictionAlgo: lfuAlgo,
		Capacity: 0,
		MaxCapacity: 10,
	}

	c.Add("1", "first cache")
	c.Add("2", "second cache")

	c.Evict()

	fifoAlgo := &cache.FifoAlgo{}
	c.SetEvictionAlgo(fifoAlgo)

	c.Add("3", "third cache")
	c.Evict()
}
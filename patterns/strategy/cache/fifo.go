package cache

import "fmt"

// FifoAlgo ...
type FifoAlgo struct {}

// Evict ...
func (fa *FifoAlgo) Evict(c *Cache) {
	fmt.Println("fifo eviction")
	c.Capacity--
}
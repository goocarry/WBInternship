package cache

import "fmt"

// LfuAlgo ...
type LfuAlgo struct {}

// Evict ...
func (la *LfuAlgo) Evict(c *Cache) {
	fmt.Println("lfu eviction")
	c.Capacity--
}
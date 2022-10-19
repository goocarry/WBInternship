package cache

import "fmt"

// LruAlgo ...
type LruAlgo struct {}

// Evict ...
func (la *LruAlgo) Evict(c *Cache) {
	fmt.Println("lru eviction")
	c.Capacity--
}
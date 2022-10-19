package cache

// EvictionAlgo ...
type EvictionAlgo interface {	
	Evict(c *Cache)
}
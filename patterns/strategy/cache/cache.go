package cache

// Cache ...
type Cache struct {
	Storage map[string]string
	EvictionAlgo EvictionAlgo
	Capacity int
	MaxCapacity int
}

// InitCache ...
func InitCache(e EvictionAlgo) *Cache {
    storage := make(map[string]string)
    return &Cache{
        Storage:      storage,
        EvictionAlgo: e,
        Capacity:     0,
        MaxCapacity:  2,
    }
}

// SetEvictionAlgo ...
func (c *Cache) SetEvictionAlgo(e EvictionAlgo) {
    c.EvictionAlgo = e
}

// Add ...
func (c *Cache) Add(key, value string) {
    if c.Capacity == c.MaxCapacity {
        c.Evict()
    }
    c.Capacity++
    c.Storage[key] = value
}

// Get ...
func (c *Cache) Get(key string) {
    delete(c.Storage, key)
}

// Evict ...
func (c *Cache) Evict() {
    c.EvictionAlgo.Evict(c)
    c.Capacity--
}
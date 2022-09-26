package cache

import (
	"sync"
	"time"

	"github.com/goocarry/wb-internship/internal/model"
)

// TODO: use interface

// Cache ...
type Cache struct {
	sync.RWMutex
	orders map[string]model.Order
}

// New ...
func New(defaultExpiration, cleanupInterval time.Duration) *Cache {

	orders := make(map[string]model.Order)

	cache := Cache{
		orders: orders,
	}

	return &cache
}

// Set ...
func (c *Cache) Set(key string, value interface{}) {

	c.Lock()

	defer c.Unlock()

	c.orders[key] = model.Order{}
}

// Get ...
func (c *Cache) Get(key string) (interface{}, bool) {

	c.RLock()

	defer c.RUnlock()

	order, found := c.orders[key]

	// key not found
	if !found {
		return nil, false
	}

	return order, true
}

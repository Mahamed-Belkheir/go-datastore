package store

import (
	"go-datastore/types"
	"sync"
)

type Cache struct {
	storage map[string]*types.TcpMessage
	mutex   sync.RWMutex
}

func (c *Cache) Get(key string) *types.TcpMessage {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.storage[key]
}

func (c *Cache) Set(key string, value *types.TcpMessage) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.storage[key] = value
}

func NewCache() *Cache {
	return &Cache{
		storage: make(map[string]*types.TcpMessage, 0),
	}
}

package store

import (
	t "go-datastore/datastructs"
	"sync"
)

type Cache struct {
	storage map[string]*t.Message
	mutex   sync.RWMutex
}

func (c *Cache) Get(key string) *t.Message {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.storage[key]
}

func (c *Cache) Set(key string, value *t.Message) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.storage[key] = value
}

func NewCache() *Cache {
	return &Cache{
		storage: make(map[string]*t.Message, 0),
	}
}

func (c *Cache) Operate(op string, data *t.Message) *t.Response {
	switch op {
	case "SET":
		c.Set(data.Key, data)
		return t.NewResponse("SET", nil)
	case "GET":
		value := c.Get(data.Key)
		if value == nil {
			return t.NewResponse("FAIL", nil)
		}
		return t.NewResponse("GET", value.Serialize())
	default:
		return t.NewResponse("FAIL", nil)
	}

}

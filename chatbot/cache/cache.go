package cache

import "sync"

type Cache struct {
	data map[string]string
	sync.Mutex
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]string),
	}
}

func (c *Cache) Get(key string) string {
	c.Lock()
	defer c.Unlock()
	return c.data[key]
}

func (c *Cache) Set(key string, value string) {
	c.Lock()
	defer c.Unlock()
	c.data[key] = value
}

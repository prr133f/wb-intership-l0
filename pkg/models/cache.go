package models

import (
	"errors"
	"sync"
)

type Cache struct {
	sync.RWMutex
	Items map[string]Item
}

type Item struct {
	Value any
}

func NewCache() *Cache {
	return &Cache{
		Items: make(map[string]Item),
	}
}

func (c *Cache) Set(key string, value any) {
	c.Lock()
	defer c.Unlock()

	c.Items[key] = Item{Value: value}
}

func (c *Cache) Get(key string) (any, bool) {
	c.RLock()
	defer c.RUnlock()

	item, found := c.Items[key]

	if !found {
		return nil, false
	}

	return item.Value, true
}

func (c *Cache) Delete(key string) error {
	c.Lock()
	defer c.Unlock()

	if _, found := c.Items[key]; !found {
		return errors.New("key not found")
	}

	delete(c.Items, key)

	return nil
}

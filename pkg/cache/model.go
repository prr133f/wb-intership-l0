package cache

import "sync"

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

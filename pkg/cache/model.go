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

type IFace interface {
	Set(key string, value any)
	Get(key string) (any, bool)
	Delete(key string) error
}

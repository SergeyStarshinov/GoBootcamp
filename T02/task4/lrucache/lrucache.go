package lrucache

import (
	"container/list"
	"sync"
)

type item[T any] struct {
	key   string
	value T
}

type Cache[T any] struct {
	capacity int
	items    map[string]*list.Element
	queue    *list.List
	mx       sync.Mutex
}

func NewMyCache[T any](capacity int) *Cache[T] {
	return &Cache[T]{
		capacity: capacity,
		items:    map[string]*list.Element{},
		queue:    list.New(),
	}
}

func (c *Cache[T]) Set(key string, value T) {
	c.mx.Lock()
	defer c.mx.Unlock()

	if element, exists := c.items[key]; exists {
		c.queue.MoveToFront(element)
		return
	}
	new := c.queue.PushFront(item[T]{key: key, value: value})
	c.items[key] = new

	if c.queue.Len() > c.capacity {
		deleted := c.queue.Remove(c.queue.Back()).(item[T])
		delete(c.items, deleted.key)
	}
}

func (c *Cache[T]) Get(key string) (T, bool) {
	c.mx.Lock()
	defer c.mx.Unlock()

	if element, exists := c.items[key]; exists {
		c.queue.MoveToFront(element)
		return element.Value.(item[T]).value, true
	}
	var value T
	return value, false
}

func (c *Cache[T]) Clear() {
	c.mx.Lock()
	defer c.mx.Unlock()

	for k := range c.items {
		delete(c.items, k)
	}
	for c.queue.Len() > 0 {
		c.queue.Remove(c.queue.Back())
	}
}

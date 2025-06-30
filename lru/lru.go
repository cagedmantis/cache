package lru

import (
	"container/list"
	"sync"
)

type entry[K comparable, T any] struct {
	key   K
	value T
}

type Cache[K comparable, T any] struct {
	mu sync.Mutex

	cap int
	l   *list.List
	m   map[K]*list.Element
}

func New[K comparable, T any](cap int) *Cache[K, T] {
	return &Cache[K, T]{
		cap: cap,
		l:   list.New(),
		m:   make(map[K]*list.Element),
	}
}

func (c *Cache[K, T]) Get(key K) (T, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	e, ok := c.m[key]
	if !ok {
		return *new(T), false
	}
	c.l.MoveToFront(e)
	val, _ := e.Value.(*entry[K, T])
	return val.value, true
}

func (c *Cache[K, T]) Put(key K, value T) {
	c.mu.Lock()
	defer c.mu.Unlock()

	e, ok := c.m[key]
	if ok {
		e.Value = &entry[K, T]{key: key, value: value}
		c.l.MoveToFront(e)
		return
	}
	e = c.l.PushFront(&entry[K, T]{key: key, value: value})
	c.m[key] = e
	if c.cap < c.l.Len() {
		re := c.l.Back()
		ent, _ := re.Value.(*entry[K, T])
		delete(c.m, ent.key)
		c.l.Remove(re)
	}
}

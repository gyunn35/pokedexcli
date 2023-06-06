package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	*sync.Mutex
	cache map[string]cacheEntry
}

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

func NewCache(t time.Duration) Cache {
	cache := Cache{
		Mutex: &sync.Mutex{},
		cache: make(map[string]cacheEntry),
	}

	go cache.reapLoop(t)

	return cache
}

func (c *Cache) Add(key string, value []byte) {
	c.Lock()
	defer c.Unlock()
	c.cache[key] = cacheEntry{
		val:       value,
		createdAt: time.Now().UTC(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.Lock()
	defer c.Unlock()
	v, ok := c.cache[key]
	return v.val, ok
}

func (c *Cache) reap(t time.Duration) {
	c.Lock()
	defer c.Unlock()

	std := time.Now().UTC().Add(-t)
	for k, v := range c.cache {
		if v.createdAt.Before(std) {
			delete(c.cache, k)
		}
	}
}

func (c *Cache) reapLoop(t time.Duration) {
	ticker := time.NewTicker(t)
	for range ticker.C {
		c.reap(t)
	}
}

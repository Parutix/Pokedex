package pokecache

import "time"

type Cache struct {
	cache map[string]cacheEntry
}

type cacheEntry struct {
	data      []byte
	createdAt time.Time
}

func NewCache(interval time.Duration) Cache {
	c := Cache {
		cache: make(map[string]cacheEntry),
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, data []byte) {
	c.cache[key] = cacheEntry {
		data: 		 data,
		createdAt: time.Now(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	entry, ok := c.cache[key]
	if !ok {
		return nil, false
	}

	return entry.data, true
}


func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval)
	}
}

func (c *Cache) reap(interval time.Duration) {
	cutoff := time.Now().Add(-interval)
	for key, value := range c.cache {
			if value.createdAt.Before(cutoff) {
					delete(c.cache, key)
			}
	}
}

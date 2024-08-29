package pokecache

import "time"

type Cache struct {
	cache map[string]cacheEntry
}

type cacheEntry struct {
	data      []byte
	createdAt time.Time
}

func NewCache() Cache {
	return Cache {
		cache: make(map[string]cacheEntry),
	}
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
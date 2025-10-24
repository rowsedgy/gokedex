package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	Entries  map[string]cacheEntry
	mu       sync.RWMutex
	interval time.Duration
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		Entries:  make(map[string]cacheEntry),
		mu:       sync.RWMutex{},
		interval: interval,
	}

	go func() {
		ticker := time.NewTicker(cache.interval)
		defer ticker.Stop()

		for range ticker.C {
			cache.reapLoop()
		}
	}()
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.Entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	entry, ok := c.Entries[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop() {
	c.mu.Lock()
	defer c.mu.Unlock()

	now := time.Now()

	for k, entry := range c.Entries {
		if now.Sub(entry.createdAt) > c.interval {
			delete(c.Entries, k)
		}
	}

}

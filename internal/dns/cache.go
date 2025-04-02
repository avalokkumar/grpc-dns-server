package dns

import (
	"sync"
	"time"
)

// Cache is a thread-safe in-memory cache for DNS resolutions
type Cache struct {
	data map[string]cacheEntry // Map to store cached entries
	mu   sync.RWMutex          // Mutex for thread-safe access
}

// cacheEntry represents a single cache entry with a value and expiration time
type cacheEntry struct {
	value     string    // Cached value (e.g., resolved IP address)
	expiresAt time.Time // Expiration time of the cache entry
}

// NewCache creates and returns a new Cache instance
func NewCache() *Cache {
	return &Cache{data: make(map[string]cacheEntry)}
}

// Get retrieves a value from the cache if it exists and is not expired
func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	// Check if the key exists in the cache
	entry, exists := c.data[key]
	if (!exists || time.Now().After(entry.expiresAt)) {
		// Return false if the key does not exist or is expired
		return "", false
	}

	// Return the cached value and true
	return entry.value, true
}

// Set adds a new key-value pair to the cache with a specified TTL
func (c *Cache) Set(key, value string, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Add the key-value pair to the cache with an expiration time
	c.data[key] = cacheEntry{value: value, expiresAt: time.Now().Add(ttl)}
}
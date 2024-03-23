package cache

import (
	"log"
	"sync"
	"time"
)

type Cache struct {
	Key        string
	Value      interface{}
	Expiration time.Time
}

type LRUCache struct {
	Size  int
	Items map[string]*Cache
	Order []string
	mutex sync.Mutex
}

func NewLRUCache(size int) *LRUCache {
	return &LRUCache{
		Size:  size,
		Items: make(map[string]*Cache),
		Order: make([]string, 0),
	}
}

// Get Retrieves the value associated with the given key from the cache
func (cache *LRUCache) Get(key string) (interface{}, bool) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	item, found := cache.Items[key]
	if !found {
		return nil, false
	}

	//Check Expiration
	if time.Now().After(item.Expiration) {
		delete(cache.Items, key)
		return nil, false
	}

	// Move on the key to front
	cache.updateOrder(key)

	return item.Value, true

}

// Set adds or updates the value associated with the given key in the cache
func (cache *LRUCache) Set(key string, value interface{}) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	// Check if key exists
	if _, found := cache.Items[key]; !found && len(cache.Items) >= cache.Size {
		// Remove least recently used item if cache is full
		cache.removeLRU()
	}

	expiration := time.Now().Add(10 * time.Second)

	// Add or update the item
	cache.Items[key] = &Cache{
		Key:        key,
		Value:      value,
		Expiration: expiration,
	}

	// Update the order
	cache.updateOrder(key)
}

// updateOrder moves the given key to the front of the order
func (cache *LRUCache) updateOrder(key string) {
	// Remove key from order if it exists
	for i, k := range cache.Order {
		if k == key {
			cache.Order = append(cache.Order[:i], cache.Order[i+1:]...)
			break
		}
	}

	// Add key to front of order
	cache.Order = append([]string{key}, cache.Order...)
	log.Println(cache.Order)
}

// removeLRU removes the least recently used item from the cache
func (cache *LRUCache) removeLRU() {
	if len(cache.Order) == 0 {
		return
	}

	lruKey := cache.Order[len(cache.Order)-1]
	delete(cache.Items, lruKey)
	cache.Order = cache.Order[:len(cache.Order)-1]
}

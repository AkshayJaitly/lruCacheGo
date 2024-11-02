package main

import (
	"container/list"
	"fmt"
)

// CacheItem represents an item in the cache with its key and value.
type CacheItem struct {
	key   string
	value string
}

// CacheError represents a custom error type for cache operations.
type CacheError struct {
	Message string
}

// Error implements the error interface for CacheError.
func (e *CacheError) Error() string {
	return e.Message
}

// LRUCache represents the Least Recently Used cache structure.
type LRUCache struct {
	capacity  int
	cacheMap  map[string]*list.Element
	cacheList *list.List
}

// NewLRUCache creates a new LRUCache with the specified capacity.
func NewLRUCache(capacity int) (*LRUCache, error) {
	if capacity <= 0 {
		return nil, &CacheError{"capacity must be greater than zero"}
	}

	return &LRUCache{
		capacity:  capacity,
		cacheMap:  make(map[string]*list.Element),
		cacheList: list.New(),
	}, nil
}

// Get retrieves the value for the specified key from the cache.
func (c *LRUCache) Get(key string) (string, error) {
	if elem, exists := c.cacheMap[key]; exists {
		c.cacheList.MoveToFront(elem)
		return elem.Value.(*CacheItem).value, nil
	}
	return "", &CacheError{fmt.Sprintf("key '%s' not found in cache", key)}
}

// Put adds a key-value pair to the cache. If the key already exists, it updates the value.
func (c *LRUCache) Put(key, value string) {
	if elem, exists := c.cacheMap[key]; exists {
		c.cacheList.MoveToFront(elem)
		elem.Value.(*CacheItem).value = value
		return
	}

	if c.cacheList.Len() == c.capacity {
		evict := c.cacheList.Back()
		c.cacheList.Remove(evict)
		delete(c.cacheMap, evict.Value.(*CacheItem).key)
	}

	item := &CacheItem{key, value}
	elem := c.cacheList.PushFront(item)
	c.cacheMap[key] = elem
}

func main() {
	cache, err := NewLRUCache(3)
	if err != nil {
		fmt.Println("Error creating cache:", err)
		return
	}

	cache.Put("1", "one")
	cache.Put("2", "two")
	cache.Put("3", "three")

	if value, err := cache.Get("1"); err == nil {
		fmt.Println(value) // should print "one"
	} else {
		fmt.Println(err)
	}

	cache.Put("4", "four")

	if value, err := cache.Get("2"); err == nil {
		fmt.Println(value) // should print empty, as it was evicted
	} else {
		fmt.Println(err) // should print "key '2' not found in cache"
	}
}

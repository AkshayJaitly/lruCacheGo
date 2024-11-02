package main

import (
	"container/list"
	"fmt"
)

type CacheItem struct {
	key   string
	value string
}

type LRUCache struct {
	capacity  int
	cacheMap  map[string]*list.Element
	cacheList *list.List
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity:  capacity,
		cacheMap:  make(map[string]*list.Element),
		cacheList: list.New(),
	}
}

func (c *LRUCache) Get(key string) (string, bool) {
	if elem, exists := c.cacheMap[key]; exists {
		c.cacheList.MoveToFront(elem)
		return elem.Value.(*CacheItem).value, true
	}
	return "", false
}

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
	cache := NewLRUCache(3)
	cache.Put("1", "one")
	cache.Put("2", "two")
	cache.Put("3", "three")
	fmt.Println(cache.Get("1")) // should be "one"
	cache.Put("4", "four")
	fmt.Println(cache.Get("2")) // should be empty, evicted
}

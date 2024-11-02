package main

import (
	"testing"
)

func TestNewLRUCache(t *testing.T) {
	cache, err := NewLRUCache(3)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if cache.capacity != 3 {
		t.Fatalf("expected capacity 3, got %d", cache.capacity)
	}
}

func TestNewLRUCacheInvalidCapacity(t *testing.T) {
	_, err := NewLRUCache(0)
	if err == nil {
		t.Fatal("expected error for capacity 0, got none")
	}
	if err.Error() != "capacity must be greater than zero" {
		t.Fatalf("expected specific error message, got %v", err)
	}
}

func TestLRUCachePutAndGet(t *testing.T) {
	cache, _ := NewLRUCache(2)

	cache.Put("1", "one")
	cache.Put("2", "two")

	if value, err := cache.Get("1"); err != nil || value != "one" {
		t.Fatalf("expected 'one', got %v with error %v", value, err)
	}

	cache.Put("3", "three") // This should evict key "2"

	if _, err := cache.Get("2"); err == nil {
		t.Fatal("expected error for getting evicted key '2', got none")
	}

	if value, err := cache.Get("3"); err != nil || value != "three" {
		t.Fatalf("expected 'three', got %v with error %v", value, err)
	}
}

func TestLRUCacheEviction(t *testing.T) {
	cache, _ := NewLRUCache(2)

	cache.Put("1", "one")
	cache.Put("2", "two")

	// Cache is full now
	cache.Put("3", "three") // This should evict key "1"

	if _, err := cache.Get("1"); err == nil {
		t.Fatal("expected error for getting evicted key '1', got none")
	}
	if value, err := cache.Get("2"); err != nil || value != "two" {
		t.Fatalf("expected 'two', got %v with error %v", value, err)
	}
	if value, err := cache.Get("3"); err != nil || value != "three" {
		t.Fatalf("expected 'three', got %v with error %v", value, err)
	}
}

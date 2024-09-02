package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Millisecond)
	if cache.cache == nil {
		t.Errorf("Expected cache to be initialized")
	}
}

func TestAddToCache(t *testing.T) {
	cache := NewCache(time.Millisecond)
	cache.Add("key", []byte("data"))
	if len(cache.cache) != 1 {
		t.Errorf("Expected cache to have 1 item")
	}
}

func TestGetFromCache(t *testing.T) {
	cache := NewCache(time.Millisecond)
	cache.Add("key", []byte("data"))
	data, ok := cache.Get("key")
	if !ok {
		t.Errorf("Expected to find key")
	}
	if string(data) != "data" {
		t.Errorf("Expected data to be 'data'")
	}
}

func TestReap(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)
	cache.Add("key", []byte("data"))
	t.Log("Added key to cache")
	
	time.Sleep(interval * 2)
	
	_, ok := cache.Get("key")
	if ok {
			t.Errorf("Expected key to be reaped")
	} else {
			t.Log("Key successfully reaped")
	}
}
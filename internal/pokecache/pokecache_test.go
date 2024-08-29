package pokecache

import "testing"

func TestCreateCache(t *testing.T) {
	cache := NewCache()
	if cache.cache == nil {
		t.Errorf("Expected cache to be initialized")
	}
}

func TestAddToCache(t *testing.T) {
	cache := NewCache()
	cache.Add("key", []byte("data"))
	if len(cache.cache) != 1 {
		t.Errorf("Expected cache to have 1 item")
	}
}

func TestGetFromCache(t *testing.T) {
	cache := NewCache()
	cache.Add("key", []byte("data"))
	data, ok := cache.Get("key")
	if !ok {
		t.Errorf("Expected to find key")
	}
	if string(data) != "data" {
		t.Errorf("Expected data to be 'data'")
	}
}
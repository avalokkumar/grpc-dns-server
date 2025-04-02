package unit

import (
	"testing"
	"time"
	"grpc-dns-server/internal/dns"
)

func TestCache(t *testing.T) {
	cache := dns.NewCache()
	cache.Set("key", "value", 1*time.Second)

	val, found := cache.Get("key")
	if !found || val != "value" {
		t.Fatalf("Expected value 'value', got '%s'", val)
	}

	time.Sleep(2 * time.Second)
	_, found = cache.Get("key")
	if found {
		t.Fatalf("Expected cache to expire, but value was found")
	}
}
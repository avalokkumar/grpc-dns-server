package integration

import (
	"testing"
	"grpc-dns-server/internal/dns"
)

func TestResolve(t *testing.T) {
	ip, err := dns.Resolve("example.com")
	if err != nil {
		t.Fatalf("Failed to resolve domain: %v", err)
	}
	if ip == "" {
		t.Fatalf("Expected a valid IP address, got empty string")
	}
}
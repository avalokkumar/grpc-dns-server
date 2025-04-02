package integration

import (
	"testing"
	"grpc-dns-server/internal/dns"
)

func TestResolveDoH(t *testing.T) {
	_, err := dns.ResolveDoH("example.com")
	if err == nil {
		t.Fatalf("Expected error for unimplemented DoH resolution, got nil")
	}
}
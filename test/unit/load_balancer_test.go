package unit

import (
	"testing"
	"grpc-dns-server/internal/dns"
)

func TestSelectServer(t *testing.T) {
	servers := []string{"8.8.8.8", "8.8.4.4"}
	selected := dns.SelectServer(servers)
	if selected != "8.8.8.8" && selected != "8.8.4.4" {
		t.Fatalf("Selected server is not in the list of servers")
	}
}
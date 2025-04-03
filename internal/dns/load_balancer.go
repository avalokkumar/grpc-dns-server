package dns

import "math/rand"

// SelectServer selects a random DNS server from the provided list
func SelectServer(servers []string) string {
	// Return a randomly selected server from the list
	return servers[rand.Intn(len(servers))]
}
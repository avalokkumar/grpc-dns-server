package dns

import "net"

// Resolve queries the upstream DNS servers to resolve the given domain name
func Resolve(domain string) (string, error) {
	// Perform a DNS lookup for the given domain
	ips, err := net.LookupIP(domain)
	if err != nil {
		// Return an error if the lookup fails
		return "", err
	}

	// Return the first resolved IP address if available
	if len(ips) > 0 {
		return ips[0].String(), nil
	}

	// Return an empty string if no IP addresses are found
	return "", nil
}
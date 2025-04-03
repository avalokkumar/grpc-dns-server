package main

import (
	"flag"
	"log"
	"grpc-dns-server/internal/grpc"
)

func main() {
	// Define the --domain flag
	domain := flag.String("domain", "", "The domain name to resolve")
	flag.Parse() // Parse the flags

	// Check if the domain flag is provided
	if *domain == "" {
		log.Fatalf("Usage: ./dns-client --domain <domain>")
	}

	// Create a new gRPC client to connect to the server at localhost:50051
	client := grpc.NewClient("localhost:50051")

	// Send a DNS resolution request for the provided domain
	response, err := client.Resolve(*domain)
	if err != nil {
		log.Fatalf("Error resolving domain: %v", err)
	}

	// Log the resolved IP address
	log.Printf("Resolved IP: %s", response)
}
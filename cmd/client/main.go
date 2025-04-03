package main

import (
	"fmt"
	"log"
	"os"
	"grpc-dns-server/internal/grpc"
)

func main() {
	// Check if a domain name is provided as an argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./dns-client <domain>")
		os.Exit(1)
	}

	// Read the domain name from command-line arguments
	domain := os.Args[1]

	// Create a new gRPC client to connect to the server at localhost:50051
	client := grpc.NewClient("localhost:50051")

	// Send a DNS resolution request for the provided domain
	response, err := client.Resolve(domain)
	if err != nil {
		log.Fatalf("Error resolving domain: %v", err)
	}

	// Log the resolved IP address
	log.Printf("Resolved IP: %s", response)
}
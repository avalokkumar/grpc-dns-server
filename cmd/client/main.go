package main

import (
	"log"
	"grpc-dns-server/internal/grpc"
)

func main() {
	// Create a new gRPC client to connect to the server at localhost:50051
	client := grpc.NewClient("localhost:50051")

	// Send a DNS resolution request for the domain "example.com"
	response, err := client.Resolve("example.com")
	if err != nil {
		log.Fatalf("Error resolving domain: %v", err)
	}

	// Log the resolved IP address
	log.Printf("Resolved IP: %s", response)
}
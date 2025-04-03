package main

import (
	"log"
	"grpc-dns-server/internal/grpc"
	"grpc-dns-server/internal/logger"
	"grpc-dns-server/pkg/utils"
)

func main() {
	// Initialize the logger for structured logging
	logger.InitLogger()

	// Load configuration from the config.yaml file
	config := utils.LoadConfig("config/config.yaml")

	// Create a new gRPC server instance with the loaded configuration
	server := grpc.NewServer(config)

	// Log the server startup message
	log.Println("Starting gRPC DNS server...")

	// Start the gRPC server
	server.Start()
}
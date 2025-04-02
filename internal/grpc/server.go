package grpc

import (
	"context"
	"log"
	"net"

	"grpc-dns-server/internal/dns"
	"grpc-dns-server/pkg/proto"
	"grpc-dns-server/pkg/utils"

	"google.golang.org/grpc"
)

// Server represents the gRPC server
type Server struct {
	grpcServer *grpc.Server // gRPC server instance
}

// dnsService implements the proto.DnsServiceServer interface
type dnsService struct {
	proto.UnimplementedDnsServiceServer // Embedding for forward compatibility
}

// Resolve handles the gRPC request to resolve a domain name
func (s *dnsService) Resolve(ctx context.Context, req *proto.ResolveRequest) (*proto.ResolveResponse, error) {
	log.Printf("Received request to resolve domain: %s", req.GetDomain())

	// Use the DNS resolver to resolve the domain
	ip, err := dns.Resolve(req.GetDomain())
	if err != nil {
		log.Printf("Failed to resolve domain: %v", err)
		return nil, err
	}

	// Return the resolved IP address
	return &proto.ResolveResponse{Ip: ip}, nil
}

// NewServer creates a new gRPC server with the provided configuration
func NewServer(config *utils.Config) *Server {
	// Use the configuration as needed (e.g., for logging, metrics, etc.)
	log.Printf("Starting server on port: %d", config.GRPC.Port)

	return &Server{grpcServer: grpc.NewServer()}
}

// Start starts the gRPC server and listens for incoming requests
func (s *Server) Start() {
	// Create a TCP listener on the specified port
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Register the DnsService with the gRPC server
	proto.RegisterDnsServiceServer(s.grpcServer, &dnsService{})

	// Start serving incoming requests
	if err := s.grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
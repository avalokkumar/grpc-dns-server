package grpc

import (
	"context"
	"log"
	"google.golang.org/grpc"
	"grpc-dns-server/pkg/proto"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conn proto.DnsServiceClient
}

func NewClient(address string) *Client {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	return &Client{conn: proto.NewDnsServiceClient(conn)}
}

func (c *Client) Resolve(domain string) (string, error) {
	resp, err := c.conn.Resolve(context.Background(), &proto.ResolveRequest{Domain: domain})
	if err != nil {
		return "", err
	}
	return resp.Ip, nil
}
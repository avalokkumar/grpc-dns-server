package utils

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func HandleError(err error) error {
	return status.Errorf(codes.Internal, "Internal server error: %v", err)
}
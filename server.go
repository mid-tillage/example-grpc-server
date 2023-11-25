package main

import (
	"context"
	"fmt"
	"log"
	"net"

    "github.com/sys-internals/example-grpc-server/userpb"
	"google.golang.org/grpc"
)

type userServiceServer struct{}

func (s *userServiceServer) GetUserByID(ctx context.Context, req *user.UserRequest) (*user.UserResponse, error) {
	// In a real application, you would fetch user data from a database based on the user ID
	// For simplicity, we'll return a dummy user here
	return &user.UserResponse{
		UserId:   req.UserId,
		Username: "john_doe",
		Email:    "john@example.com",
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	user.RegisterUserServiceServer(grpcServer, &userServiceServer{})

	fmt.Println("gRPC Server is listening on :8080...")
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

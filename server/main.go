package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/n704/go_grpc_eg"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

func (s *server) CalculateSquare(ctx context.Context, in *pb.Int32) (*pb.Int32, error) {
	newValue := in.Value * in.Value
	fmt.Printf("Square of %d is %d\n", in.Value, newValue)
	return &pb.Int32{Value: newValue}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

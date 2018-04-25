package main

import (
	"flag"
	"log"
	"time"

	pb "github.com/n704/go_grpc_eg"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = 1
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	square := flag.Int64("square", 1, "find square of number")
	cube := flag.Int64("cube", 1, "find cube of number")
	flag.Parse()
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.CalculateSquare(ctx, &pb.Int64{Value: *square})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Square of %d: %d\n", *square, r.Value)
	r, err = c.CalculateCube(ctx, &pb.Int64{Value: *cube})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Cube of %d: %d\n", *cube, r.Value)

}

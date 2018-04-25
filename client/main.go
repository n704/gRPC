package main

import (
	"log"
	"os"
	"strconv"
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

	// Contact the server and print out its response.
	value := defaultName
	if len(os.Args) > 1 {
		value, _ = strconv.Atoi(os.Args[1])

	}
	i32value := int32(value)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.CalculateSquare(ctx, &pb.Int32{Value: i32value})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Print(r.Value)
}

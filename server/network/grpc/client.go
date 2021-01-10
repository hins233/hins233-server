package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"server/server/network/grpc/pb"
	"time"
)

const (
	address = "localhost:9999"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAddServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Add(ctx, &pb.AddRequest{A: 101, B: 102})
	if err != nil {
		log.Fatalf("could not add: %v", err)
	}
	log.Printf("Sum: %d", r.GetRes())
}

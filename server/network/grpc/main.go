package main

import (
	"google.golang.org/grpc"
	"net"

	"server/server/network/grpc/pb"
)

func main() {
	createServer()
}

func createServer() {
	s, _ := net.Listen("tcp", ":9999")

	myservice := mygrpc.MyService{}

	grpcServer := grpc.NewServer()

	pb.RegisterAddServiceServer(grpcServer, &myservice)

	_ = grpcServer.Serve(s)
}

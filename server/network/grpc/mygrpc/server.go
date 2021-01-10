package mygrpc

import (
	"context"
	"server/server/network/grpc/pb"
)

type MyService struct {
}

func (s *MyService) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddReply, error) {
	res := myAdd(req.A, req.B)
	return &pb.AddReply{Res: res}, nil
}

func myAdd(a, b int32) int32 {
	return a + b
}

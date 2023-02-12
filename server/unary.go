package main

import (
	"context"

	pb "github.com/IRSHIT033/go-grpc/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *helloServer) SayHello(ctx context.Context, in *emptypb.Empty) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "hello",
	}, nil

}

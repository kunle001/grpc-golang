package main

import (
	"context"

	pb "github.com/kunle001/go-grpc/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam)(*pb.HelloResponse, error){
	 return &pb.HelloResponse{
		Message: "Hello",
	 }, nil
}
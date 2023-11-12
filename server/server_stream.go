package main

import (
	"log"
	"time"

	pb "github.com/kunle001/go-grpc/proto"
)

// emulating sending stream to a client
func(s*helloServer) SayHelloServerStreaming(req *pb.NamesList, stream pb.GreetService_SayHelloServerStreamingServer)error{
	log.Printf("got request with names: %v", req.Names)

	for _, name := range req.Names{
		res:= &pb.HelloResponse{
			Message: "Hello " + name,
		}

		if err:= stream.Send(res); err!= nil{
			return err
		}

		time.Sleep(5 *time.Second) //blocks the loop for 2 sec before it goes to the next iteration
	}

	return nil
}
package main

import (
	"context"
	"io"
	"log"

	pb "github.com/kunle001/go-grpc/proto"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NamesList){
	log.Printf("streaming has started")

	stream, err:=client.SayHelloServerStreaming(context.Background(), names);

	if err !=nil{
		log.Fatalf("could not send names: %v", err)
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF{
			break
		}
		
		if err != nil{
			log.Fatalf("error while streaming %v", err)
		}
		log.Println(message)
	}	

	log.Println("streaming ended")
}
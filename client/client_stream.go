package main

import (
	"context"
	"log"
	"time"

	pb "github.com/kunle001/go-grpc/proto"
)


func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NamesList){
	log.Printf("client streaming started")

	stream, err := client.SayHelloClientStreaming(context.Background());

	if err != nil{
		log.Fatalf("could not send names:  %v", err)
	}

	for _, name := range names.Names{
		req:= &pb.HelloRequest{
			Name: name,
		}
		if err:=stream.Send(req); err !=nil{
			log.Fatalf("Error while sending %v", err)
		}
		log.Printf("sent the request with name: %s", name)
		time.Sleep(5 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	log.Printf("client streaming finfished")

	if err!= nil{
		log.Fatalf("Eror while recieving %v", err)
	}

	log.Printf("%v", res.Messages)
}
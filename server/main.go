package main

import (
	"log"
	"net"

	pb "github.com/kunle001/go-grpc/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type helloServer struct{
	pb.GreetServiceServer
}

func main(){
	lis, err:= net.Listen("tcp", port)

	if err!=nil{
		log.Fatal("Failed to start the server", err)
	}
	
	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	log.Printf("server started at %v", lis.Addr())

	if err:= grpcServer.Serve(lis); err !=nil{
		log.Fatalf("failed to start grpc: %v", err)
	} 
}
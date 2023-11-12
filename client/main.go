package main

import (
	"log"

	pb "github.com/kunle001/go-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	// calling the server
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err!= nil{
		log.Fatalf("client couldn't connect %v", err)
	}

	defer conn.Close()

	client := pb.NewGreetServiceClient(conn);

	names := &pb.NamesList{
		Names: []string{"Adefarasin", "Abel", "Alice", "Joy", "wole"},
	}

	// callSayHello(client)
	// callSayHelloServerStream(client, names)
	callSayHelloClientStream(client, names)
}
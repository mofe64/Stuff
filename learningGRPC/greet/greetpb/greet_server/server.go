package main

import (
	"fmt"
	"google.golang.org/grpc"
	"learningGRPC/greet/greetpb"
	"log"
	"net"
)

type server struct {
}

func main() {
	fmt.Println("Hello world")

	// create listener on default port 50051
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}
	// create grpc server
	s := grpc.NewServer()

	// Register a service
	greetpb.RegisterGreetServiceServer(s, &server{})

	// bing port to grpc server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}

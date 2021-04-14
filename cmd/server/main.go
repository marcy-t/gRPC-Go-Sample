package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "grpc-golang-sample/pkg/proto"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct{}

// SayHelloメソッド
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v\n", in.Name)
	return &pb.HelloReply{
		Message: "Hello " + in.Name,
	}, nil
}

func (s *server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v\n", in.Name)
	return &pb.HelloReply{
		Message: "Hello again " + in.Name,
	}, nil
}

func main() {
	// listen
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen: %v\n", err)
	}
	fmt.Printf("Start %v\n", lis)

	// Up Server
	grpc := grpc.NewServer()
	pb.RegisterGreeterServer(grpc, &server{})

	if err := grpc.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\ns", err)
	}

}

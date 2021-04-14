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
	log.Printf("Received: %v", in.Name)
	return &pb.HelloReply{
		Message: "Hello " + in.Name,
	}, nil
}

/*
func (s *server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{
		Message: "Hello again " + in.Name,
	}, nil
}
*/

func main() {
	// listen
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	fmt.Printf("Start %v", lis)

	// Up Server
	//s := grpc.NewServer()
	grpc := grpc.NewServer()
	//pb.RegisterGreeterServer(s, &pb.UnimplementedGreeterServer{})

	pb.RegisterGreeterServer(grpc, &server{})

	/*
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	*/
	if err := grpc.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

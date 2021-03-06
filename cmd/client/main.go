package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "grpc-golang-sample/pkg/proto"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "World"
)

func main() {
	// gRPC Connection
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// args
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	// context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Call SayHello
	r1, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("Cloud not greet: %v", err)
	}

	// Call SayHelloAgain
	r2, err := c.SayHelloAgain(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("Cloud not greet SayHelloAgain: %v", err)
	}
	log.Printf("Greeting SayHello: %v\n", r1.Message)
	log.Printf("Greeting SayHelloAgain: %v\n", r2.Message)
}

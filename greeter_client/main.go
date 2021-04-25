package main

import (
	"context"
	"google.golang.org/grpc"
	pb "kj-tech.net/proto"
	"log"
	"os"
	"time"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatal("could not greet: %v", err)
	}

	log.Printf("Greeting: %s", r.GetMessage())
	r, err = c.SayHelloAgain(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatal("could not greet: %v", err)
	}

	log.Printf("Greeting: %s", r.GetMessage())
}

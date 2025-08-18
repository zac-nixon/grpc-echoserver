package main

import (
	"context"
	"log"
	"time"

	pb "echo-server/echo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewEchoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Echo(ctx, &pb.EchoRequest{Message: "Hello World2222"})
	if err != nil {
		log.Fatalf("could not echo: %v", err)
	}
	log.Printf("Echo: %s", r.GetMessage())

	r2, err := c.FixedResponse(ctx, &pb.FixedResponseRequest{})
	if err != nil {
		log.Fatalf("could not fr: %v", err)
	}
	log.Printf("fr: %s", r2.GetMessage())
}

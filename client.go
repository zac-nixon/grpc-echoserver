package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"time"

	pb "echo-server/echo"
)

func main() {
	target := "grpcaaa-1564044849.us-east-1.elb.amazonaws.com:443"
	fmt.Println(target)
	fmt.Println("using skip verify")
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true, // This skips all certificate verification, including expiry.
	}

	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)), grpc.WithAuthority("foo.com"))
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

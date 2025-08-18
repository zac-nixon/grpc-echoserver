package main

import (
	"context"
	"log"
	"net"
	"os"

	pb "echo-server/echo"
	"google.golang.org/grpc"
)

type server struct {
	fr string
	pb.UnimplementedEchoServiceServer
}

func (s *server) Echo(_ context.Context, req *pb.EchoRequest) (*pb.Response, error) {
	return &pb.Response{Message: req.Message}, nil
}

func (s *server) FixedResponse(_ context.Context, _ *pb.FixedResponseRequest) (*pb.Response, error) {
	return &pb.Response{Message: s.fr}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterEchoServiceServer(s, &server{
		fr: os.Args[1],
	})

	log.Println("gRPC server listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

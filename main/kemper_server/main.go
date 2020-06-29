package main

import (
	"context"
	"log"
	"net"

	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedDockerServerServer
}

func (s *server) ContainerInspect(ctx context.Context, in *pb.ContainerInspectRequest) (*pb.ContainerInspectReply, error) {
	return &pb.ContainerInspectReply{
		ID: in.GetID(),
		Error: &pb.ContainerError{
			Error:   "",
			Success: true,
		},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDockerServerServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

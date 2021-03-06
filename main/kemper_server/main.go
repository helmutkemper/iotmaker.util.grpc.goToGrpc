package main

import (
	"github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/server"
	"log"
	"net"

	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDockerServerServer(s, &server.GRpcServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

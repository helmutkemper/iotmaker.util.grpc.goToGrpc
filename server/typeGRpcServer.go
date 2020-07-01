package server

import (
	iotmakerDockerInterface "github.com/helmutkemper/iotmaker.docker.interface"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

type GRpcServer struct {
	dockerSystem iotmakerDockerInterface.DockerSystem
	init         bool
	pb.UnimplementedDockerServerServer
}

package server

import (
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

type GRpcServer struct {
	dockerSystem iotmakerDocker.DockerSystem
	init         bool
	pb.UnimplementedDockerServerServer
}

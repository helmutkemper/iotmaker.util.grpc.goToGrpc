package server

import (
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.0"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

type GRpcServer struct {
	dockerSystem iotmakerdocker.DockerSystem
	init         bool
	pb.UnimplementedDockerServerServer
}

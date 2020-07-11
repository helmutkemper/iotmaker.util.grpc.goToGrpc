package server

import (
	"context"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

func (el *GRpcServer) ContainerFindIdByName(
	ctx context.Context,
	in *pb.ContainerFindIdByNameRequest,
) (
	response *pb.ContainerFindIdByNameReply,
	err error,
) {

	_ = ctx
	err = el.Init()
	if err != nil {
		return
	}

	var containerID string
	err, containerID = el.dockerSystem.ContainerFindIdByName(in.GetName())

	return &pb.ContainerFindIdByNameReply{
		ID: containerID,
	}, err
}

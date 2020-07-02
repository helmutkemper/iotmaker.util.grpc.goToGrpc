package server

import (
	"context"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

func (el *GRpcServer) ContainerFindIdByNameContains(
	ctx context.Context,
	in *pb.ContainerFindIdByNameContainsRequest,
) (
	response *pb.ContainerFindIdByNameContainsReply,
	err error,
) {

	_ = ctx
	err = el.Init()
	if err != nil {
		return
	}

	var containerID string
	err, containerID = el.dockerSystem.ContainerFindIdByNameContains(in.GetName())

	return &pb.ContainerFindIdByNameContainsReply{
		ContainerID: containerID,
	}, err
}

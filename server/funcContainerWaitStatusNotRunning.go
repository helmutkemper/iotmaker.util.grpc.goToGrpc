package server

import (
	"context"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

func (el *GRpcServer) ContainerWaitStatusNotRunning(
	ctx context.Context,
	in *pb.ContainerWaitStatusNotRunningRequest,
) (
	response *pb.Empty,
	err error,
) {

	_ = ctx
	err = el.Init()
	if err != nil {
		return
	}

	err = el.dockerSystem.ContainerWaitStatusNotRunning(in.GetID())

	response = &pb.Empty{}

	return
}

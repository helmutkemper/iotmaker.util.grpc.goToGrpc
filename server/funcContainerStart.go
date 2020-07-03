package server

import (
	"context"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

func (el *GRpcServer) ContainerStart(
	ctx context.Context,
	in *pb.ContainerStartRequest,
) (
	err error,
) {

	_ = ctx
	err = el.Init()
	if err != nil {
		return
	}

	err = el.dockerSystem.ContainerStart(in.GetID())

	return
}

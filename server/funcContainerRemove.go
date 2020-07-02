package server

import (
	"context"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

func (el *GRpcServer) ContainerRemove(
	ctx context.Context,
	in *pb.ContainerRemoveRequest,
) (
	err error,
) {

	_ = ctx
	err = el.Init()
	if err != nil {
		return
	}

	err = el.dockerSystem.ContainerRemove(in.GetID(), in.GetRemoveVolumes(), in.GetRemoveLinks(), in.GetForce())

	return
}

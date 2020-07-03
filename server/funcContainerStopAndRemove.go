package server

import (
	"context"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

func (el *GRpcServer) ContainerStopAndRemove(
	ctx context.Context,
	in *pb.ContainerStopAndRemoveRequest,
) (
	response *pb.Empty,
	err error,
) {

	_ = ctx
	err = el.Init()
	if err != nil {
		return
	}

	err = el.dockerSystem.ContainerStopAndRemove(in.GetID(), in.GetRemoveVolumes(), in.GetRemoveLinks(), in.GetForce())

	response = &pb.Empty{}

	return
}

package server

import (
	"context"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

func (el *GRpcServer) ImageRemove(
	ctx context.Context,
	in *pb.ImageRemoveRequest,
) (
	response *pb.Empty,
	err error,
) {

	_ = ctx
	err = el.Init()
	if err != nil {
		return
	}

	err = el.dockerSystem.ImageRemove(in.GetID(), in.GetForce(), in.GetPruneChildren())

	response = &pb.Empty{}

	return
}

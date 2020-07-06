package server

import (
	"context"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

func (el *GRpcServer) ImageRemoveByName(
	ctx context.Context,
	in *pb.ImageRemoveByNameRequest,
) (
	response *pb.Empty,
	err error,
) {

	_ = ctx
	err = el.Init()
	if err != nil {
		return
	}

	err = el.dockerSystem.ImageRemoveByName(in.GetID(), in.GetForce(), in.GetPruneChildren())

	response = &pb.Empty{}

	return
}

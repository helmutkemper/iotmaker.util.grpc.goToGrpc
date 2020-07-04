package server

import (
	"context"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

func (el *GRpcServer) NetworkRemoveByName(
	ctx context.Context,
	in *pb.NetworkRemoveByNameRequest,
) (
	response *pb.Empty,
	err error,
) {

	_ = ctx
	err = el.Init()
	if err != nil {
		return
	}

	err = el.dockerSystem.NetworkRemoveByName(in.GetName())
	if err != nil {
		return nil, err
	}

	response = &pb.Empty{}

	return
}

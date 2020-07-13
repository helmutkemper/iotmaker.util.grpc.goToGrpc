package server

import (
	"context"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

func (el *GRpcServer) NetworkRemove(
	ctx context.Context,
	in *pb.NetworkRemoveRequest,
) (
	response *pb.Empty,
	err error,
) {

	_ = ctx
	err = el.Init()
	if err != nil {
		return
	}

	err = el.dockerSystem.NetworkRemove(in.GetID())
	if err != nil {
		return nil, err
	}

	for _, v := range networkControl {
		if in.GetID() == v.ID {
			delete(networkControl, v.Name)
			break
		}
	}

	response = &pb.Empty{}

	return
}

package server

import (
	"context"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

func (el *GRpcServer) NetworkFindIdByName(
	ctx context.Context,
	in *pb.NetworkFindIdByNameRequest,
) (
	response *pb.NetworkFindIdByNameReply,
	err error,
) {

	_ = ctx
	err = el.Init()
	if err != nil {
		return
	}

	var networkID string

	err, networkID = el.dockerSystem.NetworkFindIdByName(
		in.GetName(),
	)
	if err != nil {
		return nil, err
	}

	response = &pb.NetworkFindIdByNameReply{
		ID: networkID,
	}

	return
}

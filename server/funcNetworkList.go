package server

import (
	"context"
	"github.com/docker/docker/api/types"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

func (el *GRpcServer) NetworkList(
	ctx context.Context,
	in *pb.Empty,
) (
	response *pb.NetworkListReply,
	err error,
) {

	_ = ctx
	err = el.Init()
	if err != nil {
		return
	}

	var list []types.NetworkResource

	err, list = el.dockerSystem.NetworkList()
	if err != nil {
		return nil, err
	}

	response = &pb.NetworkListReply{
		List: SupportArrayNetworkResourceToGRpc(list),
	}

	return
}

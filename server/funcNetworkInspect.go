package server

import (
	"context"
	"github.com/docker/docker/api/types"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

func (el *GRpcServer) NetworkInspect(
	ctx context.Context,
	in *pb.NetworkInspectRequest,
) (
	response *pb.NetworkInspectReply,
	err error,
) {

	_ = ctx
	err = el.Init()
	if err != nil {
		return
	}

	var inspect types.NetworkResource

	err, inspect = el.dockerSystem.NetworkInspect(in.GetID())
	if err != nil {
		return nil, err
	}

	var ret = NetworkInspectDataConverterDockerToGRpc(inspect)

	response = &pb.NetworkInspectReply{
		NetworkResource: ret,
	}

	return
}

package server

import (
	"context"
	"github.com/docker/go-connections/nat"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

func (el *GRpcServer) ImageListExposedPortsByName(
	ctx context.Context,
	in *pb.ImageListExposedPortsByNameRequest,
) (
	response *pb.ImageListExposedPortsByNameReply,
	err error,
) {

	_ = ctx
	err = el.Init()
	if err != nil {
		return
	}

	var portList []nat.Port
	err, portList = el.dockerSystem.ImageListExposedPortsByName(in.GetName())

	response = &pb.ImageListExposedPortsByNameReply{
		List: SupportArrayNatPortToGRpc(portList),
	}

	return
}

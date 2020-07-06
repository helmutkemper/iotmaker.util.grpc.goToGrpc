package server

import (
	"context"
	"github.com/docker/go-connections/nat"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

func (el *GRpcServer) ImageListExposedPorts(
	ctx context.Context,
	in *pb.ImageListExposedPortsRequest,
) (
	response *pb.ImageListExposedPortsReply,
	err error,
) {

	_ = ctx
	err = el.Init()
	if err != nil {
		return
	}

	var portList []nat.Port
	err, portList = el.dockerSystem.ImageListExposedPorts(in.GetID())

	response = &pb.ImageListExposedPortsReply{
		List: SupportArrayNatPortToGRpc(portList),
	}

	return
}

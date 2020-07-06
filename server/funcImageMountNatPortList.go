package server

import (
	"context"
	"github.com/docker/go-connections/nat"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

func (el *GRpcServer) ImageMountNatPortList(
	ctx context.Context,
	in *pb.ImageMountNatPortListRequest,
) (
	response *pb.ImageMountNatPortListReply,
	err error,
) {

	_ = ctx
	err = el.Init()
	if err != nil {
		return
	}

	var list nat.PortMap
	err, list = el.dockerSystem.ImageMountNatPortList(in.GetID())

	response = &pb.ImageMountNatPortListReply{
		PortMap: SupportPortMapToGRpc(list),
	}

	return
}

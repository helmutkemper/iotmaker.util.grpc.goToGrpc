package server

import (
	"context"
	"github.com/docker/go-connections/nat"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

func (el *GRpcServer) ImageMountNatPortListChangeExposed(
	ctx context.Context,
	in *pb.ImageMountNatPortListChangeExposedRequest,
) (
	response *pb.ImageMountNatPortListChangeExposedReply,
	err error,
) {

	_ = ctx
	err = el.Init()
	if err != nil {
		return
	}

	var currentPortList []nat.Port
	var changeToPortList []nat.Port

	err, currentPortList = SupportGRpcArrayPortToArrayNatPot(in.GetCurrentPortList())
	if err != nil {
		return
	}

	err, changeToPortList = SupportGRpcArrayPortToArrayNatPot(in.ChangeToPortList)
	if err != nil {
		return
	}

	var list nat.PortMap
	err, list = el.dockerSystem.ImageMountNatPortListChangeExposed(in.GetImageID(), currentPortList, changeToPortList)

	response = &pb.ImageMountNatPortListChangeExposedReply{
		PortMap: SupportPortMapToGRpc(list),
	}

	return
}

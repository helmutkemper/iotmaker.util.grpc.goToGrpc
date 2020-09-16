package server

import (
	"context"
	"encoding/json"
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

	currentPortList, err = SupportGRpcArrayPortToArrayNatPot(in.GetCurrentPortList())
	if err != nil {
		return
	}

	changeToPortList, err = SupportGRpcArrayPortToArrayNatPot(in.ChangeToPortList)
	if err != nil {
		return
	}

	var list nat.PortMap
	list, err = el.dockerSystem.ImageMountNatPortListChangeExposed(in.GetID(), currentPortList, changeToPortList)

	var data []byte
	data, err = json.Marshal(&list)
	if err != nil {
		return nil, err
	}

	response = &pb.ImageMountNatPortListChangeExposedReply{
		Data: data,
	}

	return
}

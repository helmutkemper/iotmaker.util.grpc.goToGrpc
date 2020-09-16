package server

import (
	"context"
	"encoding/json"
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
	portList, err = el.dockerSystem.ImageListExposedPorts(in.GetID())

	var data []byte
	data, err = json.Marshal(&portList)
	if err != nil {
		return nil, err
	}

	response = &pb.ImageListExposedPortsReply{
		Data: data,
	}

	return
}

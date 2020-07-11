package server

import (
	"context"
	"encoding/json"
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

	var data []byte
	data, err = json.Marshal(&portList)
	if err != nil {
		return nil, err
	}

	response = &pb.ImageListExposedPortsByNameReply{
		Data: data,
	}

	return
}

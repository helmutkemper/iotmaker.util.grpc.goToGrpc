package server

import (
	"context"
	"encoding/json"
	"github.com/docker/docker/api/types"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

func (el *GRpcServer) ContainerInspectByNameContains(
	ctx context.Context,
	in *pb.ContainerInspectByNameContainsRequest,
) (
	response *pb.ContainerInspectByNameContainsReply,
	err error,
) {

	_ = ctx
	err = el.Init()
	if err != nil {
		return
	}

	var inspect []types.ContainerJSON

	inspect, err = el.dockerSystem.ContainerInspectByNameContains(in.GetName())
	if err != nil {
		return nil, err
	}

	var data []byte
	data, err = json.Marshal(&inspect)
	if err != nil {
		return nil, err
	}

	response = &pb.ContainerInspectByNameContainsReply{
		Data: data,
	}

	return
}

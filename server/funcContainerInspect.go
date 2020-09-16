package server

import (
	"context"
	"encoding/json"
	"github.com/docker/docker/api/types"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

func (el *GRpcServer) ContainerInspect(
	ctx context.Context,
	in *pb.ContainerInspectRequest,
) (
	response *pb.ContainerInspectReply,
	err error,
) {

	_ = ctx
	err = el.Init()
	if err != nil {
		return
	}

	var data []byte
	var inspect types.ContainerJSON
	inspect, err = el.dockerSystem.ContainerInspect(in.GetID())
	if err != nil {
		return nil, err
	}

	data, err = json.Marshal(&inspect)

	response = &pb.ContainerInspectReply{
		Data: data,
	}

	return
}

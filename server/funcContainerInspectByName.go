package server

import (
	"context"
	"encoding/json"
	"github.com/docker/docker/api/types"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

func (el *GRpcServer) ContainerInspectByName(
	ctx context.Context,
	in *pb.ContainerInspectByNameRequest,
) (
	response *pb.ContainerInspectByNameReply,
	err error,
) {

	_ = ctx
	err = el.Init()
	if err != nil {
		return
	}

	var inspect types.ContainerJSON

	err, inspect = el.dockerSystem.ContainerInspectByName(in.GetName())
	if err != nil {
		return nil, err
	}

	var data []byte
	data, err = json.Marshal(&inspect)
	if err != nil {
		return nil, err
	}

	response = &pb.ContainerInspectByNameReply{
		Data: data,
	}

	return
}

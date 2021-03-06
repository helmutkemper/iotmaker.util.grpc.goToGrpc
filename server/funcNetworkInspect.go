package server

import (
	"context"
	"encoding/json"
	"github.com/docker/docker/api/types"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

func (el *GRpcServer) NetworkInspect(
	ctx context.Context,
	in *pb.NetworkInspectRequest,
) (
	response *pb.NetworkInspectReply,
	err error,
) {

	_ = ctx
	err = el.Init()
	if err != nil {
		return
	}

	var inspect types.NetworkResource

	inspect, err = el.dockerSystem.NetworkInspect(in.GetID())
	if err != nil {
		return nil, err
	}

	var data []byte
	data, err = json.Marshal(&inspect)
	if err != nil {
		return nil, err
	}

	response = &pb.NetworkInspectReply{
		Data: data,
	}

	return
}

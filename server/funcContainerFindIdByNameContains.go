package server

import (
	"context"
	"encoding/json"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

func (el *GRpcServer) ContainerFindIdByNameContains(
	ctx context.Context,
	in *pb.ContainerFindIdByNameContainsRequest,
) (
	response *pb.ContainerFindIdByNameContainsReply,
	err error,
) {

	_ = ctx
	err = el.Init()
	if err != nil {
		return
	}

	var data []byte
	var list interface{}
	list, err = el.dockerSystem.ContainerFindIdByNameContains(in.GetName())

	data, err = json.Marshal(&list)
	if err != nil {
		return
	}

	return &pb.ContainerFindIdByNameContainsReply{
		Data: data,
	}, err
}

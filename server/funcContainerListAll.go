package server

import (
	"context"
	"encoding/json"
	"github.com/docker/docker/api/types"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

func (el *GRpcServer) ContainerListAll(
	ctx context.Context,
	in *pb.Empty,
) (
	response *pb.ContainerListAllReply,
	err error,
) {

	_ = ctx
	err = el.Init()
	if err != nil {
		return
	}

	var list []types.Container
	list, err = el.dockerSystem.ContainerListAll()

	var data []byte
	data, err = json.Marshal(&list)
	if err != nil {
		return nil, err
	}

	response = &pb.ContainerListAllReply{
		Data: data,
	}

	return
}

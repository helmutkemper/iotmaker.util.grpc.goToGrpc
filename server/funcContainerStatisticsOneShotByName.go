package server

import (
	"context"
	"encoding/json"
	"github.com/docker/docker/api/types"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

func (el *GRpcServer) ContainerStatisticsOneShotByName(
	ctx context.Context,
	in *pb.ContainerStatisticsOneShotByNameRequest,
) (
	response *pb.ContainerStatisticsOneShotByNameReply,
	err error,
) {

	_ = ctx
	err = el.Init()
	if err != nil {
		return
	}

	var stat types.Stats
	err, stat = el.dockerSystem.ContainerStatisticsOneShot(in.GetName())

	var data []byte
	data, err = json.Marshal(&stat)
	if err != nil {
		return nil, err
	}

	response = &pb.ContainerStatisticsOneShotByNameReply{
		Data: data,
	}

	return
}

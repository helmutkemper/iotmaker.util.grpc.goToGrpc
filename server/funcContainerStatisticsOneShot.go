package server

import (
	"context"
	"encoding/json"
	"github.com/docker/docker/api/types"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

func (el *GRpcServer) ContainerStatisticsOneShot(
	ctx context.Context,
	in *pb.ContainerStatisticsOneShotRequest,
) (
	response *pb.ContainerStatisticsOneShotReply,
	err error,
) {

	_ = ctx
	err = el.Init()
	if err != nil {
		return
	}

	var stat types.Stats
	err, stat = el.dockerSystem.ContainerStatisticsOneShot(in.GetID())

	var data []byte
	data, err = json.Marshal(&stat)
	if err != nil {
		return nil, err
	}

	response = &pb.ContainerStatisticsOneShotReply{
		Data: data,
	}

	return
}

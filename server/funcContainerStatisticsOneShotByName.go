package server

import (
	"context"
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

	response = &pb.ContainerStatisticsOneShotByNameReply{
		Statistics: SupportStatsToGRpc(stat),
	}

	return
}

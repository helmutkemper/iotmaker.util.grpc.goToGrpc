package server

import (
	"context"
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

	response = &pb.ContainerStatisticsOneShotReply{
		Statistics: SupportStatsToGRpc(stat),
	}

	return
}

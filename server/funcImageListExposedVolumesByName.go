package server

import (
	"context"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

func (el *GRpcServer) ImageListExposedVolumesByName(
	ctx context.Context,
	in *pb.ImageListExposedVolumesByNameRequest,
) (
	response *pb.ImageListExposedVolumesByNameReply,
	err error,
) {

	_ = ctx
	err = el.Init()
	if err != nil {
		return
	}

	var list []string
	list, err = el.dockerSystem.ImageListExposedVolumesByName(in.GetName())

	response = &pb.ImageListExposedVolumesByNameReply{
		List: list,
	}

	return
}

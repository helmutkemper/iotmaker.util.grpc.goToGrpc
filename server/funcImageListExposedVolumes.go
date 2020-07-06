package server

import (
	"context"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

func (el *GRpcServer) ImageListExposedVolumes(
	ctx context.Context,
	in *pb.ImageListExposedVolumesRequest,
) (
	response *pb.ImageListExposedVolumesReply,
	err error,
) {

	_ = ctx
	err = el.Init()
	if err != nil {
		return
	}

	var list []string
	err, list = el.dockerSystem.ImageListExposedVolumes(in.GetID())

	response = &pb.ImageListExposedVolumesReply{
		List: list,
	}

	return
}

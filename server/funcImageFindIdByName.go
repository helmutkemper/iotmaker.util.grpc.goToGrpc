package server

import (
	"context"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

func (el *GRpcServer) ImageFindIdByName(
	ctx context.Context,
	in *pb.ImageFindIdByNameRequest,
) (
	response *pb.ImageFindIdByNameReply,
	err error,
) {

	_ = ctx
	err = el.Init()
	if err != nil {
		return
	}

	var imageID string
	err, imageID = el.dockerSystem.ImageFindIdByName(in.GetName())
	if err != nil {
		return
	}

	response = &pb.ImageFindIdByNameReply{
		ID: imageID,
	}

	return
}

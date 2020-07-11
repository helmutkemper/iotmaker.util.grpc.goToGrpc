package server

import (
	"context"
	"encoding/json"
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

	var data []byte
	data, err = json.Marshal(&list)
	if err != nil {
		return nil, err
	}

	response = &pb.ImageListExposedVolumesReply{
		Data: data,
	}

	return
}

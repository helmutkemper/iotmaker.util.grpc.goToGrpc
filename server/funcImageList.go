package server

import (
	"context"
	"encoding/json"
	"github.com/docker/docker/api/types"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

func (el *GRpcServer) ImageList(
	ctx context.Context,
	in *pb.Empty,
) (
	response *pb.ImageListReply,
	err error,
) {

	_ = ctx
	err = el.Init()
	if err != nil {
		return
	}

	var imageList []types.ImageSummary
	imageList, err = el.dockerSystem.ImageList()

	var data []byte
	data, err = json.Marshal(&imageList)
	if err != nil {
		return nil, err
	}

	response = &pb.ImageListReply{
		Data: data,
	}

	return
}

package server

import (
	"context"
	"github.com/docker/docker/api/types"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

func (el *GRpcServer) ImageGarbageCollector(
	ctx context.Context,
	in *pb.Empty,
) (
	response *pb.Empty,
	err error,
) {

	_ = ctx
	_ = in
	err = el.Init()
	if err != nil {
		return
	}

	var list []types.ImageSummary
	err, list = el.dockerSystem.ImageList()
	for _, img := range list {
		if len(img.RepoTags) == 0 {
			continue
		}

		err = el.dockerSystem.ImageRemove(img.ID)
		if err != nil {
			return
		}
	}

	response = &pb.Empty{}

	return
}

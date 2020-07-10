package server

import (
	"context"
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
	err, imageList = el.dockerSystem.ImageList()

	var list = make([]*pb.ImageSummary, 0)
	for _, image := range imageList {
		list = append(list, &pb.ImageSummary{
			Containers:  image.Containers,
			Created:     image.Created,
			ID:          image.ID,
			Labels:      image.Labels,
			ParentID:    image.ParentID,
			RepoDigests: image.RepoDigests,
			RepoTags:    image.RepoTags,
			SharedSize:  image.SharedSize,
			Size:        image.Size,
			VirtualSize: image.VirtualSize,
		})
	}
	response = &pb.ImageListReply{
		List: list,
	}

	return
}

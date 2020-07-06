package server

import (
	"context"
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
	"github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/util"
)

func (el *GRpcServer) ImagePull(
	ctx context.Context,
	in *pb.ImagePullRequest,
) (
	response *pb.ImagePullReply,
	err error,
) {

	_ = ctx
	err = el.Init()
	if err != nil {
		return
	}

	var pullStatusChannel = make(chan iotmakerDocker.ContainerPullStatusSendToChannel, 1)
	var imageChannelID = util.RandId30()

	go func(c chan iotmakerDocker.ContainerPullStatusSendToChannel, imageChannelID string) {

		for {
			select {
			case status := <-c:
				pullStatusList[imageChannelID] = status

				if status.Closed == true {
					return
				}
			}
		}

	}(pullStatusChannel, imageChannelID)

	var imageName, imageID string

	err, imageName, imageID = el.dockerSystem.ImagePull(in.GetName(), &pullStatusChannel)

	response = &pb.ImagePullReply{
		ImageName: imageName,
		ImageID:   imageID,
	}

	return
}

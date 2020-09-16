package server

import (
	"context"
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.0"
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

	var pullStatusChannel = make(chan iotmakerdocker.ContainerPullStatusSendToChannel, 1)
	var imageChannelID = util.RandId30()

	go func(c chan iotmakerdocker.ContainerPullStatusSendToChannel, imageChannelID string) {

		for {
			select {
			case status := <-c:
				var tmp, _ = pullStatusList.Get(imageChannelID)
				tmp.Status = status
				tmp.Log += status.Stream

				pullStatusList.Set(imageChannelID, tmp)

				if status.Closed == true {
					return
				}
			}
		}

	}(pullStatusChannel, imageChannelID)

	var imageName, imageID string

	imageName, imageID, err = el.dockerSystem.ImagePull(in.GetName(), &pullStatusChannel)

	response = &pb.ImagePullReply{
		Name: imageName,
		ID:   imageID,
	}

	return
}

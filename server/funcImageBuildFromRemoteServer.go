package server

import (
	"context"
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
	"github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/util"
)

var pullStatusList map[string]iotmakerDocker.ContainerPullStatusSendToChannel

func (el *GRpcServer) ImageBuildFromRemoteServer(
	ctx context.Context,
	in *pb.ImageBuildFromRemoteServerRequest,
) (
	response *pb.ImageBuildFromRemoteServerReply,
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

	err = el.dockerSystem.ImageBuildFromRemoteServer(
		in.GetServerPath(),
		in.GetImageNewName(),
		in.GetImageTags(),
		&pullStatusChannel,
	)
	if err != nil {
		return
	}

	response = &pb.ImageBuildFromRemoteServerReply{
		ID: imageChannelID,
	}

	return
}

func init() {
	pullStatusList = make(map[string]iotmakerDocker.ContainerPullStatusSendToChannel)
}

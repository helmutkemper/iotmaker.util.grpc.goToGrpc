package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.0"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
	"github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/util"
	"time"
)

var pullStatusList PullStatusList
var pullStatusTicker = time.NewTicker(30 * time.Second * 60)

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

	var pullStatusChannel = make(chan iotmakerdocker.ContainerPullStatusSendToChannel, 1)
	var imageChannelID string

	for {
		imageChannelID = util.RandId30()
		if pullStatusList.Verify(imageChannelID) == false {
			break
		}
	}
	pullStatusList.Set(
		imageChannelID,
		BuildOrPullLog{
			Start: time.Now(),
		},
	)

	go func(c chan iotmakerdocker.ContainerPullStatusSendToChannel, imageChannelID string) {

		for {
			select {
			case status := <-c:
				var tmp, _ = pullStatusList.Get(imageChannelID)
				tmp.Status = status
				tmp.Log += status.Stream

				//fmt.Printf("build channel id: %v\n", imageChannelID)
				//fmt.Printf("%+v\n\n", tmp)
				pullStatusList.Set(imageChannelID, tmp)

				if status.Closed == true {
					fmt.Printf("build channel end\n\n\n")
					return
				}
			}
		}

	}(pullStatusChannel, imageChannelID)

	var body = in.GetData()
	var inData JSonImageBuildFromRemoteServer
	err = json.Unmarshal(body, &inData)
	if err != nil {
		err = errors.New("json unmarshal error: " + err.Error())
		return
	}

	go func(
		serverPath,
		imageName string,
		imageTags []string,
		pullStatusChannel *chan iotmakerdocker.ContainerPullStatusSendToChannel,
	) {
		_, err = el.dockerSystem.ImageBuildFromRemoteServer(
			serverPath,
			imageName,
			imageTags,
			pullStatusChannel,
		)
		if err != nil {
			var tmp, _ = pullStatusList.Get(imageChannelID)
			tmp.Log += "\nError: " + err.Error()
			pullStatusList.Set(imageChannelID, tmp)

			fmt.Printf("ImageBuildFromRemoteServer().error: %v\n", err.Error())
			return
		}
	}(
		inData.ServerPath,
		inData.ImageName,
		inData.ImageTags,
		&pullStatusChannel,
	)

	response = &pb.ImageBuildFromRemoteServerReply{
		ID: imageChannelID,
	}

	return
}

func init() {
	go func() {
		for {
			select {
			case <-pullStatusTicker.C:
				pullStatusList.TickerDeleteOldData()
			}
		}
	}()
}

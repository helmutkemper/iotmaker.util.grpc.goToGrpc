package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
	"github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/util"
	"time"
)

func (el *GRpcServer) ImageBuildAndContainerStartFromRemoteServer(
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
	var imageChannelID string
	var containerID string
	var networkConfig *network.NetworkingConfig = nil

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

	go func(c chan iotmakerDocker.ContainerPullStatusSendToChannel, imageChannelID string) {

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
	var inData JSonImageBuildAndContainerStartFromRemoteServer
	err = json.Unmarshal(body, &inData)
	if err != nil {
		err = errors.New("json unmarshal error: " + err.Error())
		return
	}

	if inData.NetworkName != "" {
		var found bool
		_, found = networkControl[inData.NetworkName]
		if found == false {
			err = errors.New("network not found. it is necessary to create a network with the name: " + inData.NetworkName)
			return
		}

		err, networkConfig = networkControl[inData.NetworkName].Generator.GetNext()
		if err != nil {
			err = errors.New("network generator error: " + err.Error())
			return
		}
	}

	go func(
		serverPath,
		imageName string,
		imageTags []string,
		containerName string,
		restartPolicy iotmakerDocker.RestartPolicy,
		mountVolumes []mount.Mount,
		containerNetwork *network.NetworkingConfig,
		currentPort []nat.Port,
		changeToPort []nat.Port,
		pullStatusChannel *chan iotmakerDocker.ContainerPullStatusSendToChannel,
	) {
		err = el.dockerSystem.ImageBuildFromRemoteServer(
			serverPath,
			imageName,
			imageTags,
			pullStatusChannel,
		)
		if err != nil {
			var tmp, _ = pullStatusList.Get(imageChannelID)
			tmp.Log += "\nError: " + err.Error()
			pullStatusList.Set(imageChannelID, tmp)

			fmt.Printf("funcImageBuildAndContainerStartFromRemoteServer().error: %v\n", err.Error())
			return
		}
		err, containerID = el.dockerSystem.ContainerCreateAndChangeExposedPort(
			imageName,
			containerName,
			restartPolicy,
			mountVolumes,
			containerNetwork,
			currentPort,
			changeToPort,
		)
		if err != nil {
			var tmp, _ = pullStatusList.Get(imageChannelID)
			tmp.Log += "\nError: " + err.Error()
			pullStatusList.Set(imageChannelID, tmp)

			fmt.Printf("funcImageBuildAndContainerStartFromRemoteServer().error: %v\n", err.Error())
			return
		}

		var tmp, _ = pullStatusList.Get(imageChannelID)
		tmp.Status.ContainerID = containerID
		tmp.Status.SuccessfullyBuildContainer = true
		pullStatusList.Set(imageChannelID, tmp)

		err = el.dockerSystem.ContainerStart(containerID)
		if err != nil {
			var tmp, _ = pullStatusList.Get(imageChannelID)
			tmp.Log += "\nError: " + err.Error()
			pullStatusList.Set(imageChannelID, tmp)

			fmt.Printf("funcImageBuildAndContainerStartFromRemoteServer().error: %v\n", err.Error())
			return
		}

	}(
		inData.ServerPath,
		inData.ImageName,
		inData.ImageTags,
		inData.ContainerName,
		inData.RestartPolicy,
		inData.MountVolumes,
		networkConfig,
		inData.CurrentPort,
		inData.ChangeToPort,
		&pullStatusChannel,
	)

	response = &pb.ImageBuildFromRemoteServerReply{
		ID: imageChannelID,
	}

	return
}

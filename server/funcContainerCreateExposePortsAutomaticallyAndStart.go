package server

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/docker/docker/api/types/network"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

func (el *GRpcServer) ContainerCreateExposePortsAutomaticallyAndStart(
	ctx context.Context,
	in *pb.ContainerCreateExposePortsAutomaticallyAndStartRequest,
) (
	response *pb.ContainerCreateExposePortsAutomaticallyAndStartReply,
	err error,
) {

	_ = ctx
	err = el.Init()
	if err != nil {
		return
	}

	var containerID string
	var body = in.GetData()
	var inData containerCreate
	var networkConfig *network.NetworkingConfig = nil

	err = json.Unmarshal(body, &inData)
	if err != nil {
		err = errors.New("json unmarshal error: " + err.Error())
		return
	}

	containerID, err = el.dockerSystem.ContainerFindIdByName(inData.ContainerName)
	if err != nil && errors.Is(err, errors.New("container name not found")) {
		err = errors.New("container find by name error: " + err.Error())
		return
	}

	if containerID != "" {
		err = errors.New("a container with this name already exists")
		return
	}

	if inData.NetworkName != "" {
		var found bool
		_, found = networkControl[inData.NetworkName]
		if found == false {
			err = errors.New("network not found. it is necessary to create a network with the name: " + inData.NetworkName)
			return
		}

		networkConfig, err = networkControl[inData.NetworkName].Generator.GetNext()
		if err != nil {
			err = errors.New("network generator error: " + err.Error())
			return
		}
	}

	containerID, err = el.dockerSystem.ContainerCreateAndExposePortsAutomatically(
		inData.ImageName,
		inData.ContainerName,
		inData.RestartPolicy,
		inData.MountVolumes,
		networkConfig,
	)

	return &pb.ContainerCreateExposePortsAutomaticallyAndStartReply{
			ID: containerID,
		},
		err
}

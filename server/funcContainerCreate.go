package server

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.0"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

type containerCreate struct {
	ImageName     string
	ContainerName string
	RestartPolicy iotmakerdocker.RestartPolicy
	MountVolumes  []mount.Mount
	NetworkName   string
	PortList      []nat.Port
}

func (el *GRpcServer) ContainerCreate(
	ctx context.Context,
	in *pb.ContainerCreateRequest,
) (
	response *pb.ContainerCreateReply,
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

	var convertedPortList = nat.PortMap{}
	for _, port := range inData.PortList {
		convertedPortList[port] = []nat.PortBinding{
			{
				HostPort: port.Port(),
			},
		}
	}

	containerID, err = el.dockerSystem.ContainerCreate(
		inData.ImageName,
		inData.ContainerName,
		inData.RestartPolicy,
		convertedPortList,
		inData.MountVolumes,
		networkConfig,
	)

	return &pb.ContainerCreateReply{
		ID: containerID,
	}, err
}

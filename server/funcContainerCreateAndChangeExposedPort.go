package server

import (
	"context"
	"encoding/json"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/go-connections/nat"
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

type containerCreateAndChangeExposedPort struct {
	ImageName        string
	ContainerName    string
	RestartPolicy    iotmakerDocker.RestartPolicy
	MountVolumes     []mount.Mount
	ContainerNetwork string
	CurrentPort      []nat.Port
	ChangeToPort     []nat.Port
}

func (el *GRpcServer) ContainerCreateAndChangeExposedPort(
	ctx context.Context,
	in *pb.ContainerCreateAndChangeExposedPortRequest,
) (
	response *pb.ContainerCreateAndChangeExposedPortReply,
	err error,
) {

	_ = ctx
	err = el.Init()
	if err != nil {
		return
	}

	var containerID string
	var body = in.GetData()
	var inData containerCreateAndChangeExposedPort

	err = json.Unmarshal(body, &inData)
	if err != nil {
		return
	}
	err, containerID = el.dockerSystem.ContainerCreateAndChangeExposedPort(
		inData.ImageName,
		inData.ContainerName,
		inData.RestartPolicy,
		inData.MountVolumes,
		nil,
		inData.CurrentPort,
		inData.ChangeToPort,
	)

	response = &pb.ContainerCreateAndChangeExposedPortReply{
		ID: containerID,
	}

	return
}

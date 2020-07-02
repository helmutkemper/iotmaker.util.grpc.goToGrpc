package server

import (
	"context"
	"github.com/docker/go-connections/nat"
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

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
	var currentPort, changeToPort []nat.Port
	var restartPolicy iotmakerDocker.RestartPolicy
	err, restartPolicy = SupportGRpcToContainerPolicy(in.GetRestartPolicy())
	if err != nil {
		return
	}

	err, currentPort = SupportGRpcArrayPortToArrayNatPot(in.GetCurrentPort())
	err, changeToPort = SupportGRpcArrayPortToArrayNatPot(in.GetChangeToPort())

	err, containerID = el.dockerSystem.ContainerCreateAndChangeExposedPort(
		in.GetImageName(),
		in.GetContainerName(),
		restartPolicy,
		SupportGRpcToArrayMount(in.GetMountVolumes()),
		nil,
		currentPort,
		changeToPort,
	)

	response = &pb.ContainerCreateAndChangeExposedPortReply{
		ContainerID: containerID,
	}

	return
}

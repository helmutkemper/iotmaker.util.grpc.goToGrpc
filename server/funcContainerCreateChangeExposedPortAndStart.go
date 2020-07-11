package server

import (
	"context"
	"github.com/docker/go-connections/nat"
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

func (el *GRpcServer) ContainerCreateChangeExposedPortAndStart(
	ctx context.Context,
	in *pb.ContainerCreateChangeExposedPortAndStartRequest,
) (
	response *pb.ContainerCreateChangeExposedPortAndStartReply,
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
	if err != nil {
		return
	}

	err, changeToPort = SupportGRpcArrayPortToArrayNatPot(in.GetChangeToPort())
	if err != nil {
		return
	}

	err, containerID = el.dockerSystem.ContainerCreateChangeExposedPortAndStart(
		in.GetImageName(),
		in.GetContainerName(),
		restartPolicy,
		SupportGRpcToArrayMount(in.GetMountVolumes()),
		nil,
		currentPort,
		changeToPort,
	)

	response = &pb.ContainerCreateChangeExposedPortAndStartReply{
		ID: containerID,
	}

	return
}

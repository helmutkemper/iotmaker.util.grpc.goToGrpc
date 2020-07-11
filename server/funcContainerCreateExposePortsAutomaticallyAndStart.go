package server

import (
	"context"
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
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
	var restartPolicy iotmakerDocker.RestartPolicy
	err, restartPolicy = SupportGRpcToContainerPolicy(in.GetRestartPolicy())
	if err != nil {
		return
	}

	err, containerID = el.dockerSystem.ContainerCreateExposePortsAutomaticallyAndStart(
		in.GetImageName(),
		in.GetContainerName(),
		restartPolicy,
		SupportGRpcToArrayMount(in.GetMountVolumes()),
		nil,
	)

	response = &pb.ContainerCreateExposePortsAutomaticallyAndStartReply{
		ID: containerID,
	}

	return
}

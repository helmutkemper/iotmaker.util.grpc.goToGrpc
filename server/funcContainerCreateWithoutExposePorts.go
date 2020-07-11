package server

import (
	"context"
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

func (el *GRpcServer) ContainerCreateWithoutExposePorts(
	ctx context.Context,
	in *pb.ContainerCreateWithoutExposePortsRequest,
) (
	response *pb.ContainerCreateWithoutExposePortsReply,
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

	err, containerID = el.dockerSystem.ContainerCreateWithoutExposePorts(
		in.GetImageName(),
		in.GetContainerName(),
		restartPolicy,
		SupportGRpcToArrayMount(in.GetMountVolumes()),
		nil,
	)

	return &pb.ContainerCreateWithoutExposePortsReply{
		ID: containerID,
	}, err
}

package server

import (
	"context"
	"github.com/docker/go-connections/nat"
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

func (el *GRpcServer) ContainerCreateAndStart(
	ctx context.Context,
	in *pb.ContainerCreateAndStartRequest,
) (
	response *pb.ContainerCreateAndStartReply,
	err error,
) {

	_ = ctx
	err = el.Init()
	if err != nil {
		return
	}

	var containerID string
	var restartPolicy iotmakerDocker.RestartPolicy
	var portExposedList nat.PortMap
	err, restartPolicy = SupportGRpcToContainerPolicy(in.GetRestartPolicy())
	if err != nil {
		return
	}

	err, portExposedList = SupportGRpcToNatPotMap(in.GetPortExposedList())
	if err != nil {
		return
	}

	err, containerID = el.dockerSystem.ContainerCreateAndStart(
		in.GetImageName(),
		in.GetContainerName(),
		restartPolicy,
		portExposedList,
		SupportGRpcToArrayMount(in.GetMountVolumes()),
		nil,
	)

	return &pb.ContainerCreateAndStartReply{
		ID: containerID,
	}, err
}

package server

import (
	"context"
	"github.com/docker/docker/api/types"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

func (el *GRpcServer) ContainerInspect(
	ctx context.Context,
	in *pb.ContainerInspectRequest,
) (
	response *pb.ContainerInspectReply,
	err error,
) {

	_ = ctx
	err = el.Init()
	if err != nil {
		return
	}

	var inspect types.ContainerJSON

	err, inspect = el.dockerSystem.ContainerInspect(in.GetID())
	if err != nil {
		return nil, err
	}

	var ret = ContainerInspectDataConverterDockerToGRpc(inspect)

	response = &pb.ContainerInspectReply{
		ID:            in.GetID(),
		ContainerJSON: ret,
	}

	return
}

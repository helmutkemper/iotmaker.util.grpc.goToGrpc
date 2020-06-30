package server

import (
	"context"
	"github.com/docker/docker/api/types"
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

type GRpcServer struct {
	pb.UnimplementedDockerServerServer
}

func (s *GRpcServer) NetworkInspect(
	ctx context.Context,
	in *pb.NetworkInspectRequest,
) (
	response *pb.NetworkInspectReply,
	err error,
) {

	_ = ctx
	var inspect types.NetworkResource

	d := iotmakerDocker.DockerSystem{}
	err = d.Init()
	if err != nil {
		return &pb.NetworkInspectReply{
			ID: in.GetID(),
		}, err
	}

	err, inspect = d.NetworkInspect(in.GetID())
	if err != nil {
		return &pb.NetworkInspectReply{
			ID: in.GetID(),
		}, err
	}

	var ret = NetworkInspectDataConverterDockerToGRpc(inspect)

	response = &pb.NetworkInspectReply{
		ID:              in.GetID(),
		NetworkResource: ret,
	}

	return
}

func (s *GRpcServer) ContainerInspect(
	ctx context.Context,
	in *pb.ContainerInspectRequest,
) (
	response *pb.ContainerInspectReply,
	err error,
) {

	_ = ctx
	var inspect types.ContainerJSON

	d := iotmakerDocker.DockerSystem{}
	err = d.Init()
	if err != nil {
		return &pb.ContainerInspectReply{
			ID: in.GetID(),
		}, err
	}
	err, inspect = d.ContainerInspect(in.GetID())
	if err != nil {
		return &pb.ContainerInspectReply{
			ID: in.GetID(),
		}, err
	}

	var ret = ContainerInspectDataConverterDockerToGRpc(inspect)

	response = &pb.ContainerInspectReply{
		ID:            in.GetID(),
		ContainerJSON: ret,
	}

	return
}

package server

import (
	"context"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

func (el *GRpcServer) ContainerInspectJSonByName(
	ctx context.Context,
	in *pb.ContainerInspectJSonByNameRequest,
) (
	response *pb.ContainerInspectJSonByNameReply,
	err error,
) {

	_ = ctx
	err = el.Init()
	if err != nil {
		return
	}

	var inspect []byte

	err, inspect = el.dockerSystem.ContainerInspectJSonByName(in.GetName())
	if err != nil {
		return nil, err
	}

	response = &pb.ContainerInspectJSonByNameReply{
		Inspect: string(inspect),
	}

	return
}

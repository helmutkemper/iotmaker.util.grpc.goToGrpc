package server

import (
	"context"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

func (el *GRpcServer) ContainerInspectJSon(
	ctx context.Context,
	in *pb.ContainerInspectJSonRequest,
) (
	response *pb.ContainerInspectJSonReply,
	err error,
) {

	_ = ctx
	err = el.Init()
	if err != nil {
		return
	}

	var inspect []byte

	err, inspect = el.dockerSystem.ContainerInspectJSon(in.GetID())
	if err != nil {
		return nil, err
	}

	response = &pb.ContainerInspectJSonReply{
		Inspect: string(inspect),
	}

	return
}

package server

import (
	"context"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

func (el *GRpcServer) ImageGarbageCollector(
	ctx context.Context,
	in *pb.Empty,
) (
	response *pb.Empty,
	err error,
) {

	_ = ctx
	_ = in
	err = el.Init()
	if err != nil {
		return
	}

	err = el.dockerSystem.ImageGarbageCollector()

	response = &pb.Empty{}

	return
}

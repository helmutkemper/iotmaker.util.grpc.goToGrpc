package server

import (
	"context"
	"errors"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

func (el *GRpcServer) ImageBuildFromRemoteServerStatus(
	ctx context.Context,
	in *pb.ImageOrContainerBuildPullStatusRequest,
) (
	response *pb.ImageOrContainerBuildPullStatusReply,
	err error,
) {

	_ = ctx

	status, found := pullStatusList[in.GetID()]
	if found == false {
		err = errors.New("image build id not found")
	}

	response = SupportBuildStatusToGRpc(status)

	return
}

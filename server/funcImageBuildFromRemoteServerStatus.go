package server

import (
	"context"
	"encoding/json"
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

	status, found := pullStatusList.Get(in.GetID())
	if found == false {
		err = errors.New("image build id not found")
		return
	}

	var data []byte
	data, err = json.Marshal(&status)
	if err != nil {
		return nil, err
	}

	response = &pb.ImageOrContainerBuildPullStatusReply{
		Data: data,
	}

	return
}

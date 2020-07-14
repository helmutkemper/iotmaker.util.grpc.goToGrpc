package server

import (
	"context"
	"encoding/json"
	"errors"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

func (el *GRpcServer) ContainerStart(
	ctx context.Context,
	in *pb.ContainerStartRequest,
) (
	response *pb.Empty,
	err error,
) {

	_ = ctx
	err = el.Init()
	if err != nil {
		return
	}

	var body = in.GetData()
	var inData JSonContainerRemove
	err = json.Unmarshal(body, &inData)
	if err != nil {
		err = errors.New("json unmarshal error: " + err.Error())
		return
	}

	err = el.dockerSystem.ContainerStart(inData.Id)
	response = &pb.Empty{}

	return
}

package server

import (
	"context"
	"encoding/json"
	"errors"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

type JSonContainerGenericRequest struct {
	Name    string
	Id      string
	Volumes bool
	Links   bool
	Force   bool
}

func (el *GRpcServer) ContainerRemove(
	ctx context.Context,
	in *pb.ContainerRemoveRequest,
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
	var inData JSonContainerGenericRequest
	err = json.Unmarshal(body, &inData)
	if err != nil {
		err = errors.New("json unmarshal error: " + err.Error())
		return
	}

	err = el.dockerSystem.ContainerRemove(
		inData.Id,
		inData.Volumes,
		inData.Links,
		inData.Force,
	)

	response = &pb.Empty{}

	return
}

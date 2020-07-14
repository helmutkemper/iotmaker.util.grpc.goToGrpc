package server

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/docker/docker/api/types"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

func (el *GRpcServer) ContainerStatisticsOneShotByName(
	ctx context.Context,
	in *pb.ContainerStatisticsOneShotByNameRequest,
) (
	response *pb.ContainerStatisticsOneShotReply,
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

	var stat types.Stats
	err, stat = el.dockerSystem.ContainerStatisticsOneShotByName(inData.Name)

	var data []byte
	data, err = json.Marshal(&stat)
	if err != nil {
		return nil, err
	}

	response = &pb.ContainerStatisticsOneShotReply{
		Data: data,
	}

	return
}

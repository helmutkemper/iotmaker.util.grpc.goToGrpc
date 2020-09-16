package server

import (
	"github.com/docker/go-connections/nat"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

func SupportGRpcArrayPortToArrayNatPot(
	pt []*pb.Port,
) (
	ret []nat.Port,
	err error,
) {

	ret = make([]nat.Port, 0)
	for _, p := range pt {
		var port nat.Port
		port, err = nat.NewPort(p.Protocol, p.Port)
		if err != nil {
			return
		}
		ret = append(ret, port)
	}

	return
}

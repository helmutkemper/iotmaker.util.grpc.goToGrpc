package server

import (
	"github.com/docker/go-connections/nat"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

func SupportArrayNatPortToGRpc(list []nat.Port) (pt []*pb.Port) {
	pt = make([]*pb.Port, 0)
	for _, port := range list {
		pt = append(pt, &pb.Port{
			Port:     port.Port(),
			Protocol: port.Proto(),
		})
	}

	return
}

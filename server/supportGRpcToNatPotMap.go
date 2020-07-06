package server

import (
	"github.com/docker/go-connections/nat"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

func SupportGRpcToNatPotMap(
	pt *pb.PortMap,
) (
	err error,
	portMap nat.PortMap,
) {

	portMap = nat.PortMap{}

	var port nat.Port
	for portString, v := range pt.Port {
		err, port = SupportStringToPort(portString)
		if err != nil {
			return
		}

		var toBind = make([]nat.PortBinding, 0)
		for _, b := range v.PortBinding {
			toBind = append(toBind, nat.PortBinding{
				HostIP:   b.HostIP,
				HostPort: b.HostPort,
			})
		}

		portMap[port] = toBind
	}

	return
}

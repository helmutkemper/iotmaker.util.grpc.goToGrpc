package server

import (
	"github.com/docker/go-connections/nat"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

func SupportPortMapToGRpc(portMap nat.PortMap) (pt *pb.PortMap) {
	var port = make(map[string]*pb.PortBindingList)
	pt = &pb.PortMap{}

	for p, list := range portMap {
		var portList = make([]*pb.PortBinding, 0)
		for _, v := range list {
			portList = append(portList, &pb.PortBinding{
				HostIP:   v.HostIP,
				HostPort: v.HostPort,
			})
		}
		port[string(p)] = &pb.PortBindingList{
			PortBinding: portList,
		}
	}
	pt.Port = port

	return
}

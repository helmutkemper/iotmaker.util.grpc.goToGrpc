package server

import (
	"context"
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

func (el *GRpcServer) NetworkCreate(
	ctx context.Context,
	in *pb.NetworkCreateRequest,
) (
	response *pb.NetworkCreateReply,
	err error,
) {

	_ = ctx
	err = el.Init()
	if err != nil {
		return
	}

	var networkID string
	var networkGenerator *iotmakerDocker.NextNetworkAutoConfiguration

	err, networkID, networkGenerator = el.dockerSystem.NetworkCreate(
		in.GetName(),
		SupportStringToNetworkDrive(in.GetNetworkDrive()),
		in.GetScope(),
		in.GetSubnet(),
		in.GetGateway(),
	)
	if err != nil {
		return nil, err
	}

	if len(networkControl) == 0 {
		networkControl = make(map[string]NetworkControl)
	}
	networkControl[networkID] = NetworkControl{
		Generator: networkGenerator,
		Name:      in.GetName(),
		Drive:     in.GetNetworkDrive(),
		Scope:     in.GetScope(),
		Subnet:    in.GetSubnet(),
		Gateway:   in.GetGateway(),
	}

	response = &pb.NetworkCreateReply{
		NetworkID: networkID,
	}

	return
}

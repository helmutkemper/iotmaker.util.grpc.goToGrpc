package server

import (
	"context"
	"errors"
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.0"
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
	var networkGenerator *iotmakerdocker.NextNetworkAutoConfiguration
	var found bool

	_, found = networkControl[in.GetName()]
	if found == true {
		err = errors.New("there is already a network with the name: " + in.GetName())
		return
	}

	networkID, networkGenerator, err = el.dockerSystem.NetworkCreate(
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
	networkControl[in.GetName()] = NetworkControl{
		Generator: networkGenerator,
		ID:        networkID,
		Name:      in.GetName(),
		Drive:     in.GetNetworkDrive(),
		Scope:     in.GetScope(),
		Subnet:    in.GetSubnet(),
		Gateway:   in.GetGateway(),
	}

	response = &pb.NetworkCreateReply{
		ID: networkID,
	}

	return
}

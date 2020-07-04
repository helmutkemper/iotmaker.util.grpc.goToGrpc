package server

import (
	"github.com/docker/docker/api/types/network"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

func SupportNetworkEndpointSettingsToGRpc(
	endpoint *network.EndpointSettings,
) (
	settings *pb.EndpointSettings,
) {

	var endpointIPAMConfig *pb.EndpointIPAMConfig
	if endpoint.IPAMConfig != nil {
		endpointIPAMConfig = &pb.EndpointIPAMConfig{
			IPv4Address:  endpoint.IPAMConfig.IPv4Address,
			IPv6Address:  endpoint.IPAMConfig.IPv6Address,
			LinkLocalIPs: endpoint.IPAMConfig.LinkLocalIPs,
		}
	}

	settings = &pb.EndpointSettings{
		IPAMConfig:          endpointIPAMConfig,
		Links:               endpoint.Links,
		Aliases:             endpoint.Aliases,
		NetworkID:           endpoint.NetworkID,
		EndpointID:          endpoint.EndpointID,
		Gateway:             endpoint.Gateway,
		IPAddress:           endpoint.IPAddress,
		IPPrefixLen:         int64(endpoint.IPPrefixLen),
		IPv6Gateway:         endpoint.IPv6Gateway,
		GlobalIPv6Address:   endpoint.GlobalIPv6Address,
		GlobalIPv6PrefixLen: int64(endpoint.GlobalIPv6PrefixLen),
		MacAddress:          endpoint.MacAddress,
		DriverOpts:          endpoint.DriverOpts,
	}

	return
}

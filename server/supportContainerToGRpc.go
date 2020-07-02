package server

import (
	"github.com/docker/docker/api/types"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

func SupportContainerToGRpc(cont []types.Container) (protoContainers []*pb.Container) {
	protoContainers = make([]*pb.Container, 0)

	for _, container := range cont {
		var containerPorts = make([]*pb.ContainerPort, 0)
		for _, p := range container.Ports {
			containerPorts = append(containerPorts, &pb.ContainerPort{
				IP:          p.IP,
				PrivatePort: uint32(p.PrivatePort),
				PublicPort:  uint32(p.PublicPort),
				Type:        p.Type,
			})
		}

		var containerHostConfig = &pb.ContainerHostConfig{
			NetworkMode: container.HostConfig.NetworkMode,
		}

		var containerNetworks = make(map[string]*pb.EndpointSettings)
		if container.NetworkSettings != nil {
			for k, endpoint := range container.NetworkSettings.Networks {
				var endpointIPAMConfig *pb.EndpointIPAMConfig
				if endpoint.IPAMConfig != nil {
					endpointIPAMConfig = &pb.EndpointIPAMConfig{
						IPv4Address:  endpoint.IPAMConfig.IPv4Address,
						IPv6Address:  endpoint.IPAMConfig.IPv6Address,
						LinkLocalIPs: endpoint.IPAMConfig.LinkLocalIPs,
					}
				}

				containerNetworks[k] = &pb.EndpointSettings{
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
			}
		}

		var containerNetworkSettings = &pb.SummaryNetworkSettings{
			Networks: containerNetworks,
		}
		protoContainers = append(protoContainers, &pb.Container{
			ID:      container.ID,
			Names:   container.Names,
			Image:   container.Image,
			ImageID: container.ImageID,
			Command: container.Command,
			Created: container.Created,

			Ports: containerPorts,

			SizeRw:     container.SizeRw,
			SizeRootFs: container.SizeRootFs,
			Labels:     container.Labels,
			State:      container.State,
			Status:     container.Status,

			HostConfig:      containerHostConfig,
			NetworkSettings: containerNetworkSettings,

			Mounts: SupportTypeMountToGRpc(container.Mounts),
		})
	}

	return
}

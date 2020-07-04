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
				containerNetworks[k] = SupportNetworkEndpointSettingsToGRpc(endpoint)
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

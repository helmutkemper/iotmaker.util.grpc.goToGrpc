package server

import (
	"github.com/docker/docker/api/types"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

func SupportArrayNetworkResourceToGRpc(
	listIn []types.NetworkResource,
) (
	listOut []*pb.NetworkResource,
) {

	listOut = make([]*pb.NetworkResource, 0)

	for _, resource := range listIn {
		var resourceIPAMConfig = make([]*pb.IPAMConfig, 0)
		for _, conf := range resource.IPAM.Config {
			resourceIPAMConfig = append(resourceIPAMConfig, &pb.IPAMConfig{
				Subnet:     conf.Subnet,
				IPRange:    conf.IPRange,
				Gateway:    conf.Gateway,
				AuxAddress: conf.AuxAddress,
			})
		}

		var resourceIPAm = &pb.IPAM{
			Driver:  resource.IPAM.Driver,
			Options: resource.IPAM.Options,
			Config:  resourceIPAMConfig,
		}

		var resourceConfigFrom = &pb.ConfigReference{
			Network: resource.ConfigFrom.Network,
		}

		var resourceContainers = make(map[string]*pb.EndpointResource)
		for k, container := range resource.Containers {
			resourceContainers[k] = &pb.EndpointResource{
				Name:        container.Name,
				EndpointID:  container.EndpointID,
				MacAddress:  container.MacAddress,
				IPv4Address: container.IPv4Address,
				IPv6Address: container.IPv6Address,
			}
		}

		var resourcePeers = make([]*pb.PeerInfo, 0)
		for _, info := range resource.Peers {
			resourcePeers = append(resourcePeers, &pb.PeerInfo{
				Name: info.Name,
				IP:   info.IP,
			})
		}

		var resourceServices = make(map[string]*pb.ServiceInfo)
		for k, info := range resource.Services {
			var infoTasks = make([]*pb.Task, 0)
			for _, info := range info.Tasks {
				infoTasks = append(infoTasks, &pb.Task{
					Name:       info.Name,
					EndpointID: info.EndpointID,
					EndpointIP: info.EndpointIP,
					Info:       info.Info,
				})
			}
			resourceServices[k] = &pb.ServiceInfo{
				VIP:          info.VIP,
				Ports:        info.Ports,
				LocalLBIndex: int64(info.LocalLBIndex),
				Tasks:        infoTasks,
			}
		}

		listOut = append(listOut, &pb.NetworkResource{
			Name:       resource.Name,
			ID:         resource.ID,
			Created:    resource.Created.Unix(),
			Scope:      resource.Scope,
			Driver:     resource.Driver,
			EnableIPv6: resource.EnableIPv6,

			IPAM: resourceIPAm,

			Internal:   resource.Internal,
			Attachable: resource.Attachable,
			Ingress:    resource.Ingress,

			ConfigFrom: resourceConfigFrom,

			ConfigOnly: resource.ConfigOnly,

			Containers: resourceContainers,

			Options: resource.Options,
			Labels:  resource.Labels,

			Peers:    resourcePeers,
			Services: resourceServices,
		})
	}

	return
}

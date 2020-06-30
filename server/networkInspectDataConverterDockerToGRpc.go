package server

import (
	"github.com/docker/docker/api/types"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

func NetworkInspectDataConverterDockerToGRpc(
	data types.NetworkResource,
) (
	ret *pb.NetworkResource,
) {

	var containers = make(map[string]*pb.EndpointResource)
	for k, resource := range data.Containers {
		var endPointResource = &pb.EndpointResource{
			Name:        resource.Name,
			EndpointID:  resource.EndpointID,
			MacAddress:  resource.MacAddress,
			IPv4Address: resource.IPv4Address,
			IPv6Address: resource.IPv4Address,
		}
		containers[k] = endPointResource
	}

	var services = make(map[string]*pb.ServiceInfo)
	for k, service := range data.Services {
		var task = make([]*pb.Task, 0)
		for _, taskValue := range service.Tasks {
			task = append(task, &pb.Task{
				Name:       taskValue.Name,
				EndpointID: taskValue.EndpointID,
				EndpointIP: taskValue.EndpointIP,
				Info:       taskValue.Info,
			})
		}

		var info = &pb.ServiceInfo{
			VIP:          service.VIP,
			Ports:        service.Ports,
			LocalLBIndex: int64(service.LocalLBIndex),
			Tasks:        task,
		}

		services[k] = info
	}

	var dataPeers = make([]*pb.PeerInfo, 0)
	for _, peer := range data.Peers {
		dataPeers = append(dataPeers, &pb.PeerInfo{
			Name: peer.Name,
			IP:   peer.IP,
		})
	}

	var config = make([]*pb.IPAMConfig, 0)
	for _, ipConfig := range data.IPAM.Config {
		config = append(config, &pb.IPAMConfig{
			Subnet:     ipConfig.Subnet,
			IPRange:    ipConfig.IPRange,
			Gateway:    ipConfig.Gateway,
			AuxAddress: ipConfig.AuxAddress,
		})
	}

	ret = &pb.NetworkResource{
		Name:       data.Name,
		ID:         data.ID,
		Created:    data.Created.Unix(),
		Scope:      data.Scope,
		Driver:     data.Driver,
		EnableIPv6: data.EnableIPv6,

		IPAM: &pb.IPAM{
			Driver:  data.IPAM.Driver,
			Options: data.IPAM.Options,
			Config:  config,
		},

		Internal:   data.Internal,
		Attachable: data.Attachable,
		Ingress:    data.Ingress,

		ConfigFrom: &pb.ConfigReference{
			Network: data.ConfigFrom.Network,
		},

		ConfigOnly: data.ConfigOnly,

		Containers: containers,

		Options: data.Options,
		Labels:  data.Labels,

		Peers: dataPeers,

		Services: services,
	}

	return
}

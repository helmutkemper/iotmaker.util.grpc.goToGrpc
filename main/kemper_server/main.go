package main

import (
	"context"
	"github.com/docker/docker/api/types"
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
	"log"
	"net"

	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedDockerServerServer
}

func ContainerInspectDataConverterDockerToGRpc(data types.ContainerJSON) (ret *pb.ContainerJSON) {

	var dataState *pb.ContainerState
	var dataNode *pb.ContainerNode
	var dataHostConfig *pb.HostConfig
	var dataResourcesHostConfigMemorySwappiness int64
	var dataHostConfigResourcesOomKillDisable bool
	var dataHostConfigResourcesPidsLimit int64
	var dataHostConfigResourcesULimits = make([]*pb.Ulimit, 0)
	var dataHostConfigResourcesBlkioDeviceReadBps = make([]*pb.ThrottleDevice, 0)
	var dataHostConfigResourcesBlkioDeviceWriteBps = make([]*pb.ThrottleDevice, 0)
	var dataHostConfigResourcesBlkioDeviceReadIOps = make([]*pb.ThrottleDevice, 0)
	var dataHostConfigResourcesBlkioDeviceWriteIOps = make([]*pb.ThrottleDevice, 0)
	var dataHostConfigResourcesDevices = make([]*pb.DeviceMapping, 0)
	var dataHostConfigResourcesDeviceRequests = make([]*pb.DeviceRequest, 0)
	var dataMounts = make([]*pb.MountPoint, 0)

	for _, point := range data.Mounts {
		dataMounts = append(dataMounts, &pb.MountPoint{
			Type:        string(point.Type),
			Name:        point.Name,
			Source:      point.Source,
			Destination: point.Destination,
			Driver:      point.Driver,
			Mode:        point.Mode,
			RW:          point.RW,
			Propagation: string(point.Propagation),
		})
	}

	var portMap = &pb.PortMap{}
	portMap.Port = make(map[string]*pb.PortBindingList)

	var dataStateHealthStatus string
	var dataStateHealthFailingStreak int64

	var dataHealthCheckResult = make([]*pb.HealthcheckResult, 0)
	if data.State != nil && data.State.Health != nil {
		for _, healthcheck := range data.State.Health.Log {
			dataHealthCheckResult = append(dataHealthCheckResult, &pb.HealthcheckResult{
				Start:    healthcheck.Start.Unix(),
				End:      healthcheck.End.Unix(),
				ExitCode: int64(healthcheck.ExitCode),
				Output:   healthcheck.Output,
			})
		}

		dataStateHealthStatus = data.State.Health.Status
		dataStateHealthFailingStreak = int64(data.State.Health.FailingStreak)
	}

	if data.State != nil {
		dataState = &pb.ContainerState{
			Status:     data.State.Status,
			Running:    data.State.Running,
			Paused:     data.State.Paused,
			Restarting: data.State.Restarting,
			OOMKilled:  data.State.OOMKilled,
			Dead:       data.State.Dead,
			Pid:        int64(data.State.Pid),
			ExitCode:   int64(data.State.ExitCode),
			Error:      data.State.Error,
			StartedAt:  data.State.StartedAt,
			FinishedAt: data.State.FinishedAt,

			Health: &pb.Health{
				Status:        dataStateHealthStatus,
				FailingStreak: dataStateHealthFailingStreak,
				Log:           dataHealthCheckResult,
			},
		}
	}

	weightDevice := make([]*pb.WeightDevice, 0)
	if data.HostConfig != nil {
		for _, weight := range data.HostConfig.Resources.BlkioWeightDevice {
			weightDevice = append(weightDevice, &pb.WeightDevice{
				Path:   weight.Path,
				Weight: uint32(weight.Weight),
			})
		}
	}

	if data.HostConfig != nil {
		if data.HostConfig.Resources.MemorySwappiness != nil {
			dataResourcesHostConfigMemorySwappiness = *data.HostConfig.Resources.MemorySwappiness
		}

		if data.HostConfig.Resources.OomKillDisable != nil {
			dataHostConfigResourcesOomKillDisable = *data.HostConfig.Resources.OomKillDisable
		}

		if data.HostConfig.Resources.PidsLimit != nil {
			dataHostConfigResourcesPidsLimit = *data.HostConfig.Resources.PidsLimit
		}

		for _, ULimits := range data.HostConfig.Resources.Ulimits {
			dataHostConfigResourcesULimits = append(dataHostConfigResourcesULimits, &pb.Ulimit{
				Name: ULimits.Name,
				Hard: ULimits.Hard,
				Soft: ULimits.Soft,
			})
		}

		for _, blkioDeviceReadBps := range data.HostConfig.Resources.BlkioDeviceReadBps {
			dataHostConfigResourcesBlkioDeviceReadBps = append(dataHostConfigResourcesBlkioDeviceReadBps, &pb.ThrottleDevice{
				Path: blkioDeviceReadBps.Path,
				Rate: blkioDeviceReadBps.Rate,
			})
		}

		for _, blkioDeviceWriteBps := range data.HostConfig.Resources.BlkioDeviceWriteBps {
			dataHostConfigResourcesBlkioDeviceReadBps = append(dataHostConfigResourcesBlkioDeviceReadBps, &pb.ThrottleDevice{
				Path: blkioDeviceWriteBps.Path,
				Rate: blkioDeviceWriteBps.Rate,
			})
		}

		for _, blkioDeviceReadIOps := range data.HostConfig.Resources.BlkioDeviceReadIOps {
			dataHostConfigResourcesBlkioDeviceReadBps = append(dataHostConfigResourcesBlkioDeviceReadBps, &pb.ThrottleDevice{
				Path: blkioDeviceReadIOps.Path,
				Rate: blkioDeviceReadIOps.Rate,
			})
		}

		for _, blkioDeviceWriteIOps := range data.HostConfig.Resources.BlkioDeviceWriteIOps {
			dataHostConfigResourcesBlkioDeviceReadBps = append(dataHostConfigResourcesBlkioDeviceReadBps, &pb.ThrottleDevice{
				Path: blkioDeviceWriteIOps.Path,
				Rate: blkioDeviceWriteIOps.Rate,
			})
		}

		for _, devices := range data.HostConfig.Resources.Devices {
			dataHostConfigResourcesDevices = append(dataHostConfigResourcesDevices, &pb.DeviceMapping{
				PathOnHost:        devices.PathOnHost,
				PathInContainer:   devices.PathInContainer,
				CgroupPermissions: devices.CgroupPermissions,
			})
		}

		for _, deviceRequests := range data.HostConfig.Resources.DeviceRequests {

			var Capabilities = make([]*pb.CapabilitiesList, 0)
			for _, capab := range deviceRequests.Capabilities {
				Capabilities = append(Capabilities, &pb.CapabilitiesList{
					Capabilities: capab,
				})
			}

			dataHostConfigResourcesDeviceRequests = append(dataHostConfigResourcesDeviceRequests, &pb.DeviceRequest{
				Driver:       deviceRequests.Driver,
				Count:        int64(deviceRequests.Count),
				DeviceIDs:    deviceRequests.DeviceIDs,
				Capabilities: Capabilities,
				Options:      deviceRequests.Options,
			})
		}

		for port, arrPortBinding := range data.HostConfig.PortBindings {
			var portBind = make([]*pb.PortBinding, 0)
			var toAdd = &pb.PortBindingList{}
			for _, bind := range arrPortBinding {
				portBind = append(portBind, &pb.PortBinding{
					HostIP:   bind.HostIP,
					HostPort: bind.HostPort,
				})
			}
			toAdd.PortBinding = portBind
			portMap.Port[string(port)] = toAdd
		}
	}

	if data.HostConfig != nil {
		dataHostConfig = &pb.HostConfig{
			Binds:           data.HostConfig.Binds,
			ContainerIDFile: data.HostConfig.ContainerIDFile,
			LogConfig: &pb.LogConfig{
				Type:   data.HostConfig.LogConfig.Type,
				Config: data.HostConfig.LogConfig.Config,
			},
			NetworkMode:  string(data.HostConfig.NetworkMode),
			PortBindings: portMap,
			RestartPolicy: &pb.RestartPolicy{
				Name:              data.HostConfig.RestartPolicy.Name,
				MaximumRetryCount: int64(data.HostConfig.RestartPolicy.MaximumRetryCount),
			},
			AutoRemove:      data.HostConfig.AutoRemove,
			VolumeDriver:    data.HostConfig.VolumeDriver,
			VolumesFrom:     data.HostConfig.VolumesFrom,
			CapAdd:          data.HostConfig.CapAdd,
			CapDrop:         data.HostConfig.CapDrop,
			Capabilities:    data.HostConfig.Capabilities,
			DNS:             data.HostConfig.DNS,
			DNSOptions:      data.HostConfig.DNSOptions,
			DNSSearch:       data.HostConfig.DNSSearch,
			ExtraHosts:      data.HostConfig.ExtraHosts,
			GroupAdd:        data.HostConfig.GroupAdd,
			IpcMode:         string(data.HostConfig.IpcMode),
			Cgroup:          string(data.HostConfig.Cgroup),
			Links:           data.HostConfig.Links,
			OomScoreAdj:     int64(data.HostConfig.OomScoreAdj),
			PidMode:         string(data.HostConfig.PidMode),
			Privileged:      data.HostConfig.Privileged,
			PublishAllPorts: data.HostConfig.PublishAllPorts,
			ReadonlyRootfs:  data.HostConfig.ReadonlyRootfs,
			SecurityOpt:     data.HostConfig.SecurityOpt,
			StorageOpt:      data.HostConfig.StorageOpt,
			Tmpfs:           data.HostConfig.Tmpfs,
			UTSMode:         string(data.HostConfig.UTSMode),
			UsernsMode:      string(data.HostConfig.UsernsMode),
			ShmSize:         data.HostConfig.ShmSize,
			Sysctls:         data.HostConfig.Sysctls,
			Runtime:         data.HostConfig.Runtime,
			Isolation:       string(data.HostConfig.Isolation),

			Resources: &pb.Resources{
				CPUShares:         data.HostConfig.Resources.CPUShares,
				Memory:            data.HostConfig.Resources.Memory,
				NanoCPUs:          data.HostConfig.Resources.NanoCPUs,
				CgroupParent:      data.HostConfig.Resources.CgroupParent,
				BlkioWeight:       uint32(data.HostConfig.Resources.BlkioWeight),
				BlkioWeightDevice: weightDevice,

				BlkioDeviceReadBps:   dataHostConfigResourcesBlkioDeviceReadBps,
				BlkioDeviceWriteBps:  dataHostConfigResourcesBlkioDeviceWriteBps,
				BlkioDeviceReadIOps:  dataHostConfigResourcesBlkioDeviceReadIOps,
				BlkioDeviceWriteIOps: dataHostConfigResourcesBlkioDeviceWriteIOps,

				CPUPeriod:          data.HostConfig.Resources.CPUPeriod,
				CPUQuota:           data.HostConfig.Resources.CPUQuota,
				CPURealtimePeriod:  data.HostConfig.Resources.CPURealtimePeriod,
				CPURealtimeRuntime: data.HostConfig.Resources.CPURealtimeRuntime,
				CpusetCpus:         data.HostConfig.Resources.CpusetCpus,
				CpusetMems:         data.HostConfig.Resources.CpusetMems,

				Devices: dataHostConfigResourcesDevices,

				DeviceCgroupRules: data.HostConfig.Resources.DeviceCgroupRules,

				DeviceRequests: dataHostConfigResourcesDeviceRequests,

				KernelMemory:      data.HostConfig.Resources.KernelMemory,
				KernelMemoryTCP:   data.HostConfig.Resources.KernelMemoryTCP,
				MemoryReservation: data.HostConfig.Resources.MemoryReservation,
				MemorySwap:        data.HostConfig.Resources.MemorySwap,

				MemorySwappiness: dataResourcesHostConfigMemorySwappiness,
				OomKillDisable:   dataHostConfigResourcesOomKillDisable,
				PidsLimit:        dataHostConfigResourcesPidsLimit,
				Ulimits:          dataHostConfigResourcesULimits,

				CPUCount:           data.HostConfig.Resources.CPUCount,
				CPUPercent:         data.HostConfig.Resources.CPUPercent,
				IOMaximumIOps:      data.HostConfig.Resources.IOMaximumIOps,
				IOMaximumBandwidth: data.HostConfig.Resources.IOMaximumBandwidth,
			},

			Mounts: dataMounts,
		}
	}

	if data.Node != nil {
		dataNode = &pb.ContainerNode{
			ID:        data.Node.ID,
			IPAddress: data.Node.IPAddress,
			Addr:      data.Node.Addr,
			Name:      data.Node.Name,
			Cpus:      int64(data.Node.Cpus),
			Memory:    data.Node.Memory,
			Labels:    data.Node.Labels,
		}
	}

	ret = &pb.ContainerJSON{
		ContainerJSONBase: &pb.ContainerJSONBase{
			ID:      data.ID,
			Created: data.Created,
			Path:    data.Path,
			Args:    data.Args,

			State: dataState,

			Image:          data.Image,
			ResolvConfPath: data.ResolvConfPath,
			HostnamePath:   data.HostnamePath,
			HostsPath:      data.HostsPath,
			LogPath:        data.LogPath,

			Node: dataNode,

			Name:            data.Name,
			RestartCount:    int64(data.RestartCount),
			Driver:          data.Driver,
			Platform:        data.Platform,
			MountLabel:      data.MountLabel,
			ProcessLabel:    data.ProcessLabel,
			AppArmorProfile: data.AppArmorProfile,
			ExecIDs:         data.ExecIDs,

			HostConfig: dataHostConfig,
		},
	}

	return
}

func (s *server) NetworkInspect(
	ctx context.Context,
	in *pb.NetworkInspectRequest,
) (
	response *pb.NetworkInspectReply,
	err error,
) {

	var inspect types.NetworkResource

	d := iotmakerDocker.DockerSystem{}
	err = d.Init()
	if err != nil {
		return &pb.NetworkInspectReply{
			ID: in.GetID(),
		}, err
	}

	err, inspect = d.NetworkInspect(in.GetID())
	if err != nil {
		return &pb.NetworkInspectReply{
			ID: in.GetID(),
		}, err
	}

	var ret = NetworkInspectDataConverterDockerToGRpc(inspect)

	response = &pb.NetworkInspectReply{
		ID:              in.GetID(),
		NetworkResource: ret,
	}

	return
}

func (s *server) ContainerInspect(
	ctx context.Context,
	in *pb.ContainerInspectRequest,
) (
	response *pb.ContainerInspectReply,
	err error,
) {

	var inspect types.ContainerJSON

	d := iotmakerDocker.DockerSystem{}
	err = d.Init()
	if err != nil {
		return &pb.ContainerInspectReply{
			ID: in.GetID(),
		}, err
	}
	err, inspect = d.ContainerInspect(in.GetID())
	if err != nil {
		return &pb.ContainerInspectReply{
			ID: in.GetID(),
		}, err
	}

	var ret = ContainerInspectDataConverterDockerToGRpc(inspect)

	response = &pb.ContainerInspectReply{
		ID:            in.GetID(),
		ContainerJSON: ret,
	}

	return
}

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

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDockerServerServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

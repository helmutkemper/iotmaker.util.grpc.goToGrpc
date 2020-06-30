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

type js pb.ContainerJSON

func (el *js) FromContainer(data types.ContainerJSON) (ret *pb.ContainerJSON) {

	var dataState *pb.ContainerState
	var dataNode *pb.ContainerNode
	var dataHostConfig *pb.HostConfig
	var dataResourcesHostConfigMemorySwappiness int64
	var dataHostConfigResourcesOomKillDisable bool
	var dataHostConfigResourcesPidsLimit int64
	var dataHostConfigResourcesULimits = make([]*pb.Ulimit, 0)
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
			PortBindings: &pb.PortMap{
				//Port: data.HostConfig.PortBindings //todo: fazer
			},
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
				//BlkioDeviceReadBps:   data.HostConfig.Resources.BlkioDeviceReadBps,
				//BlkioDeviceWriteBps:  data.HostConfig.Resources.BlkioDeviceWriteBps,
				//BlkioDeviceReadIOps:  data.HostConfig.Resources.BlkioDeviceReadIOps,
				//BlkioDeviceWriteIOps: data.HostConfig.Resources.BlkioDeviceWriteIOps,
				CPUPeriod:          data.HostConfig.Resources.CPUPeriod,
				CPUQuota:           data.HostConfig.Resources.CPUQuota,
				CPURealtimePeriod:  data.HostConfig.Resources.CPURealtimePeriod,
				CPURealtimeRuntime: data.HostConfig.Resources.CPURealtimeRuntime,
				CpusetCpus:         data.HostConfig.Resources.CpusetCpus,
				CpusetMems:         data.HostConfig.Resources.CpusetMems,
				//Devices:              data.HostConfig.Resources.Devices,
				DeviceCgroupRules: data.HostConfig.Resources.DeviceCgroupRules,
				//DeviceRequests:       data.HostConfig.Resources.DeviceRequests,
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

			Mounts:        nil,
			MaskedPaths:   nil,
			ReadonlyPaths: nil,
			Init:          false,
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

			//HostConfig       : data.HostConfig,
			//GraphDriver      : data.GraphDriver,
			//SizeRw           : data.SizeRw,
			//SizeRootFs       : data.SizeRootFs,
		},
	}

	return
}

func (s *server) ContainerInspect(ctx context.Context, in *pb.ContainerInspectRequest) (*pb.ContainerInspectReply, error) {

	var err error
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

	var j js
	var ret = j.FromContainer(inspect)

	r := &pb.ContainerInspectReply{
		ID:            in.GetID(),
		ContainerJSON: ret,
	}

	return r, nil
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
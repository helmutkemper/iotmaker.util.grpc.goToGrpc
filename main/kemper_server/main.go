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

	healthcheckResult := make([]*pb.HealthcheckResult, 0)
	for _, healthcheck := range data.State.Health.Log {
		healthcheckResult = append(healthcheckResult, &pb.HealthcheckResult{
			Start:    healthcheck.Start.Unix(),
			End:      healthcheck.End.Unix(),
			ExitCode: int64(healthcheck.ExitCode),
			Output:   healthcheck.Output,
		})
	}

	ret = &pb.ContainerJSON{
		ContainerJSONBase: &pb.ContainerJSONBase{
			ID:      data.ID,
			Created: data.Created,
			Path:    data.Path,
			Args:    data.Args,

			State: &pb.ContainerState{
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
					Status:        data.State.Health.Status,
					FailingStreak: int64(data.State.Health.FailingStreak),
					Log:           healthcheckResult,
				},
			},

			Image:          data.Image,
			ResolvConfPath: data.ResolvConfPath,
			HostnamePath:   data.HostnamePath,
			HostsPath:      data.HostsPath,
			LogPath:        data.LogPath,

			Node: &pb.ContainerNode{
				ID:        data.Node.ID,
				IPAddress: data.Node.IPAddress,
				Addr:      data.Node.Addr,
				Name:      data.Node.Name,
				Cpus:      int64(data.Node.Cpus),
				Memory:    data.Node.Memory,
				Labels:    data.Node.Labels,
			},

			Name:            data.Name,
			RestartCount:    int64(data.RestartCount),
			Driver:          data.Driver,
			Platform:        data.Platform,
			MountLabel:      data.MountLabel,
			ProcessLabel:    data.ProcessLabel,
			AppArmorProfile: data.AppArmorProfile,
			ExecIDs:         data.ExecIDs,

			HostConfig: &pb.HostConfig{
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
				ShmSize:         0,
				Sysctls:         nil,
				Runtime:         "",
				Isolation:       "",
				Resources:       nil,
				Mounts:          nil,
				MaskedPaths:     nil,
				ReadonlyPaths:   nil,
				Init:            false,
			},

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

package iotmaker_util_grpc_goToGrpc

import "github.com/docker/docker/api/types"

type ContainerJSON struct {
	ContainerJSONBase
	Mounts          []MountPoint
	Config          Config
	NetworkSettings NetworkSettings
}

func (el *ContainerJSON) FromContainer(data types.ContainerJSON) {
	el.ID = data.ID
	el.Created = data.Created
	el.Path = data.Path
	el.Args = data.Args

	el.State.Status = data.State.Status
	el.State.Running = data.State.Running
	el.State.Paused = data.State.Paused
	el.State.Restarting = data.State.Restarting
	el.State.OOMKilled = data.State.OOMKilled
	el.State.Dead = data.State.Dead
	el.State.Pid = data.State.Pid
	el.State.ExitCode = data.State.ExitCode
	el.State.Error = data.State.Error
	el.State.StartedAt = data.State.StartedAt
	el.State.FinishedAt = data.State.FinishedAt
	//el.State.Health      = data.State.Health

	el.Image = data.Image
	el.ResolvConfPath = data.ResolvConfPath
	el.HostnamePath = data.HostnamePath
	el.HostsPath = data.HostsPath
	el.LogPath = data.LogPath
	//el.Node              = data.Node
	el.Name = data.Name
	el.RestartCount = data.RestartCount
	el.Driver = data.Driver
	el.Platform = data.Platform
	el.MountLabel = data.MountLabel
	el.ProcessLabel = data.ProcessLabel
	el.AppArmorProfile = data.AppArmorProfile
	el.ExecIDs = data.ExecIDs
	//el.HostConfig        = data.HostConfig
	//el.GraphDriver       = data.GraphDriver
	//el.SizeRw            = data.SizeRw
	//el.SizeRootFs        = data.SizeRootFs
}

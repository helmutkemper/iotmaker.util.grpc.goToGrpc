package server

import (
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/go-connections/nat"
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.0"
)

type JSonImageBuildAndContainerStartFromRemoteServer struct {
	ImageName     string
	ImageTags     []string
	ServerPath    string
	ContainerName string
	RestartPolicy iotmakerdocker.RestartPolicy
	MountVolumes  []mount.Mount
	NetworkName   string
	CurrentPort   []nat.Port
	ChangeToPort  []nat.Port
}

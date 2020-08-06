package server

import (
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/go-connections/nat"
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
)

type JSonImageBuildAndContainerStartFromRemoteServer struct {
	ImageName     string
	ImageTags     []string
	ServerPath    string
	ContainerName string
	RestartPolicy iotmakerDocker.RestartPolicy
	MountVolumes  []mount.Mount
	NetworkName   string
	CurrentPort   []nat.Port
	ChangeToPort  []nat.Port
}
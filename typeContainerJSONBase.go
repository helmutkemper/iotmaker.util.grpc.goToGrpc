package iotmaker_util_grpc_goToGrpc

type ContainerJSONBase struct {
	ID              string
	Created         string
	Path            string
	Args            []string
	State           ContainerState
	Image           string
	ResolvConfPath  string
	HostnamePath    string
	HostsPath       string
	LogPath         string
	Node            ContainerNode
	Name            string
	RestartCount    int
	Driver          string
	Platform        string
	MountLabel      string
	ProcessLabel    string
	AppArmorProfile string
	ExecIDs         []string
	HostConfig      HostConfig
	GraphDriver     GraphDriverData
	SizeRw          int64
	SizeRootFs      int64
}

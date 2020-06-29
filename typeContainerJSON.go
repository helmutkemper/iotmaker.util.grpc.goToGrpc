package iotmaker_util_grpc_goToGrpc

type ContainerJSON struct {
	ContainerJSONBase
	Mounts          []MountPoint
	Config          Config
	NetworkSettings NetworkSettings
}

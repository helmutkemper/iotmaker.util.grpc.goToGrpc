package iotmaker_util_grpc_goToGrpc

type MountPoint struct {
	Type        Type
	Name        string
	Source      string
	Destination string
	Driver      string
	Mode        string
	RW          bool
	Propagation Propagation
}

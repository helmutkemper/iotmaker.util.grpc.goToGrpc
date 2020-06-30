package iotmaker_util_grpc_goToGrpc

type IPAM struct {
	Driver  string
	Options map[string]string
	Config  []IPAMConfig
}

package iotmaker_util_grpc_goToGrpc

type IPAMConfig struct {
	Subnet     string
	IPRange    string
	Gateway    string
	AuxAddress map[string]string
}

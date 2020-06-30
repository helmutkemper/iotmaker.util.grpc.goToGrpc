package iotmaker_util_grpc_goToGrpc

type ServiceInfo struct {
	VIP          string
	Ports        []string
	LocalLBIndex int
	Tasks        []Task
}

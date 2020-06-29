package iotmaker_util_grpc_goToGrpc

type DeviceRequest struct {
	Driver       string
	Count        int
	DeviceIDs    []string
	Capabilities [][]string
	Options      map[string]string
}

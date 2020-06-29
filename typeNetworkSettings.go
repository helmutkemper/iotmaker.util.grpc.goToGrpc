package iotmaker_util_grpc_goToGrpc

type NetworkSettings struct {
	NetworkSettingsBase
	DefaultNetworkSettings
	Networks map[string]EndpointSettings
}

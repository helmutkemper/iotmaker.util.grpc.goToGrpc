package iotmaker_util_grpc_goToGrpc

type NetworkResource struct {
	Name       string
	ID         string
	Created    int64
	Scope      string
	Driver     string
	EnableIPv6 bool
	IPAM       IPAM
	Internal   bool
	Attachable bool
	Ingress    bool
	ConfigFrom ConfigReference
	ConfigOnly bool
	Containers map[string]EndpointResource
	Options    map[string]string
	Labels     map[string]string
	Peers      PeerInfo
	Services   map[string]ServiceInfo
}

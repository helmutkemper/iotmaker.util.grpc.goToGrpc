package server

type EndpointIPAMConfig struct {
	IPv4Address  string
	IPv6Address  string
	LinkLocalIPs []string
}

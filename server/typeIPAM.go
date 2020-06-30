package server

type IPAM struct {
	Driver  string
	Options map[string]string
	Config  []IPAMConfig
}

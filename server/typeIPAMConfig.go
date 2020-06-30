package server

type IPAMConfig struct {
	Subnet     string
	IPRange    string
	Gateway    string
	AuxAddress map[string]string
}

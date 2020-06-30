package server

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

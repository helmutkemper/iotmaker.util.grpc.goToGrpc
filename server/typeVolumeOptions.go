package server

type VolumeOptions struct {
	NoCopy       bool
	Labels       map[string]string
	DriverConfig Driver
}

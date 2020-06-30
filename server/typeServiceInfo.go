package server

type ServiceInfo struct {
	VIP          string
	Ports        []string
	LocalLBIndex int
	Tasks        []Task
}

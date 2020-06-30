package server

type Task struct {
	Name       string
	EndpointID string
	EndpointIP string
	Info       map[string]string
}

package server

type ContainerNode struct {
	ID        string
	IPAddress string
	Addr      string
	Name      string
	Cpus      int
	Memory    int64
	Labels    map[string]string
}

package server

type HostConfig struct {
	Binds           []string
	ContainerIDFile string
	LogConfig       LogConfig
	NetworkMode     NetworkMode
	PortBindings    PortMap
	RestartPolicy   RestartPolicy
	AutoRemove      bool
	VolumeDriver    string
	VolumesFrom     []string
	CapAdd          StrSlice
	CapDrop         StrSlice
	Capabilities    []string
	DNS             []string
	DNSOptions      []string
	DNSSearch       []string
	ExtraHosts      []string
	GroupAdd        []string
	IpcMode         IpcMode
	Cgroup          CgroupSpec
	Links           []string
	OomScoreAdj     int
	PidMode         PidMode
	Privileged      bool
	PublishAllPorts bool
	ReadonlyRootfs  bool
	SecurityOpt     []string
	StorageOpt      map[string]string
	Tmpfs           map[string]string
	UTSMode         UTSMode
	UsernsMode      UsernsMode
	ShmSize         int64
	Sysctls         map[string]string
	Runtime         string
	ConsoleSize     [2]uint
	Isolation       Isolation
	Resources
	Mounts        []Mount
	MaskedPaths   []string
	ReadonlyPaths []string
	Init          bool
}

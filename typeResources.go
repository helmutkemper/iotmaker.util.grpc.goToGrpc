package iotmaker_util_grpc_goToGrpc

type Resources struct {
	CPUShares            int64
	Memory               int64
	NanoCPUs             int64
	CgroupParent         string
	BlkioWeight          uint16
	BlkioWeightDevice    []WeightDevice
	BlkioDeviceReadBps   []ThrottleDevice
	BlkioDeviceWriteBps  []ThrottleDevice
	BlkioDeviceReadIOps  []ThrottleDevice
	BlkioDeviceWriteIOps []ThrottleDevice
	CPUPeriod            int64
	CPUQuota             int64
	CPURealtimePeriod    int64
	CPURealtimeRuntime   int64
	CpusetCpus           string
	CpusetMems           string
	Devices              []DeviceMapping
	DeviceCgroupRules    []string
	DeviceRequests       []DeviceRequest
	KernelMemory         int64
	KernelMemoryTCP      int64
	MemoryReservation    int64
	MemorySwap           int64
	MemorySwappiness     int64
	OomKillDisable       bool
	PidsLimit            int64
	Ulimits              []Ulimit
	CPUCount             int64
	CPUPercent           int64
	IOMaximumIOps        uint64
	IOMaximumBandwidth   uint64
}

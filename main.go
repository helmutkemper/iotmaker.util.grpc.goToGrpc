//protoc --go_opt=paths=source_relative --go_out=plugins=grpc:. *.proto
package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
)

type ContainerState struct {
	Status     string // String representation of the container state. Can be one of "created", "running", "paused", "restarting", "removing", "exited", or "dead"
	Running    bool
	Paused     bool
	Restarting bool
	OOMKilled  bool
	Dead       bool
	Pid        int
	ExitCode   int
	Error      string
	StartedAt  string
	FinishedAt string
	Health     Health `json:",omitempty"`
}

type Health struct {
	Status        string              // Status is one of Starting, Healthy or Unhealthy
	FailingStreak int                 // FailingStreak is the number of consecutive failures
	Log           []HealthcheckResult // Log contains the last few results (oldest first)
}

type zone struct {
	name   string // abbreviated name, "CET"
	offset int    // seconds east of UTC
	isDST  bool   // is this zone Daylight Savings Time?
}

type zoneTrans struct {
	when         int64 // transition time, in seconds since 1970 GMT
	index        uint8 // the index of the zone that goes into effect at that time
	isstd, isutc bool  // ignored - no idea what these mean
}

type Location struct {
	name string
	zone []zone
	tx   []zoneTrans

	// Most lookups will be for the current time.
	// To avoid the binary search through tx, keep a
	// static one-element cache that gives the correct
	// zone for the time when the Location was created.
	// if cacheStart <= t < cacheEnd,
	// lookup can return cacheZone.
	// The units for cacheStart and cacheEnd are seconds
	// since January 1, 1970 UTC, to match the argument
	// to lookup.
	cacheStart int64
	cacheEnd   int64
	cacheZone  zone
}

type Time struct {
	wall uint64
	ext  int64
	loc  Location
}

type HealthcheckResult struct {
	Start    Time   // Start is the time this check started
	End      Time   // End is the time this check ended
	ExitCode int    // ExitCode meanings: 0=healthy, 1=unhealthy, 2=reserved (considered unhealthy), else=error running probe
	Output   string // Output from last check
}

// DeviceRequest represents a request for devices from a device driver.
// Used by GPU device drivers.
type DeviceRequest struct {
	Driver       string            // Name of device driver
	Count        int               // Number of devices to request (-1 = All)
	DeviceIDs    []string          // List of device IDs as recognizable by the device driver
	Capabilities [][]string        // An OR list of AND lists of device capabilities (e.g. "gpu")
	Options      map[string]string // Options to pass onto the device driver
}

// DeviceMapping represents the device mapping between the host and the container.
type DeviceMapping struct {
	PathOnHost        string
	PathInContainer   string
	CgroupPermissions string
}

// WeightDevice is a structure that holds device:weight pair
type WeightDevice struct {
	Path   string
	Weight uint16
}

// ThrottleDevice is a structure that holds device:rate_per_second pair
type ThrottleDevice struct {
	Path string
	Rate uint64
}

// Ulimit is a human friendly version of Rlimit.
type Ulimit struct {
	Name string
	Hard int64
	Soft int64
}

// Resources contains container's resources (cgroups config, ulimits...)
type Resources struct {
	// Applicable to all platforms
	CPUShares int64 `json:"CpuShares"` // CPU shares (relative weight vs. other containers)
	Memory    int64 // Memory limit (in bytes)
	NanoCPUs  int64 `json:"NanoCpus"` // CPU quota in units of 10<sup>-9</sup> CPUs.

	// Applicable to UNIX platforms
	CgroupParent         string // Parent cgroup.
	BlkioWeight          uint16 // Block IO weight (relative weight vs. other containers)
	BlkioWeightDevice    []WeightDevice
	BlkioDeviceReadBps   []ThrottleDevice
	BlkioDeviceWriteBps  []ThrottleDevice
	BlkioDeviceReadIOps  []ThrottleDevice
	BlkioDeviceWriteIOps []ThrottleDevice
	CPUPeriod            int64           `json:"CpuPeriod"`          // CPU CFS (Completely Fair Scheduler) period
	CPUQuota             int64           `json:"CpuQuota"`           // CPU CFS (Completely Fair Scheduler) quota
	CPURealtimePeriod    int64           `json:"CpuRealtimePeriod"`  // CPU real-time period
	CPURealtimeRuntime   int64           `json:"CpuRealtimeRuntime"` // CPU real-time runtime
	CpusetCpus           string          // CpusetCpus 0-2, 0,1
	CpusetMems           string          // CpusetMems 0-2, 0,1
	Devices              []DeviceMapping // List of devices to map inside the container
	DeviceCgroupRules    []string        // List of rule to be added to the device cgroup
	DeviceRequests       []DeviceRequest // List of device requests for device drivers
	KernelMemory         int64           // Kernel memory limit (in bytes)
	KernelMemoryTCP      int64           // Hard limit for kernel TCP buffer memory (in bytes)
	MemoryReservation    int64           // Memory soft limit (in bytes)
	MemorySwap           int64           // Total memory usage (memory + swap); set `-1` to enable unlimited swap
	MemorySwappiness     int64           // Tuning container memory swappiness behaviour
	OomKillDisable       bool            // Whether to disable OOM Killer or not
	PidsLimit            int64           // Setting PIDs limit for a container; Set `0` or `-1` for unlimited, or `null` to not change.
	Ulimits              []Ulimit        // List of ulimits to be set in the container

	// Applicable to Windows
	CPUCount           int64  `json:"CpuCount"`   // CPU count
	CPUPercent         int64  `json:"CpuPercent"` // CPU percent
	IOMaximumIOps      uint64 // Maximum IOps for the container system drive
	IOMaximumBandwidth uint64 // Maximum IO in bytes per second for the container system drive
}

// Isolation represents the isolation technology of a container. The supported
// values are platform specific
type Isolation string

// UsernsMode represents userns mode in the container.
type UsernsMode string

// UTSMode represents the UTS namespace of the container.
type UTSMode string

// PidMode represents the pid namespace of the container.
type PidMode string

// CgroupSpec represents the cgroup to use for the container.
type CgroupSpec string

// IpcMode represents the container ipc stack.
type IpcMode string

// RestartPolicy represents the restart policies of the container.
type RestartPolicy struct {
	Name              string
	MaximumRetryCount int
}

type NetworkMode string

// LogConfig represents the logging configuration of the container.
type LogConfig struct {
	Type   string
	Config map[string]string
}

type StrSlice []string

type Consistency string

// BindOptions defines options specific to mounts of type "bind".
type BindOptions struct {
	Propagation  Propagation `json:",omitempty"`
	NonRecursive bool        `json:",omitempty"`
}

// Driver represents a volume driver.
type Driver struct {
	Name    string            `json:",omitempty"`
	Options map[string]string `json:",omitempty"`
}

// VolumeOptions represents the options for a mount of type volume.
type VolumeOptions struct {
	NoCopy       bool              `json:",omitempty"`
	Labels       map[string]string `json:",omitempty"`
	DriverConfig Driver            `json:",omitempty"`
}

// TmpfsOptions defines options specific to mounts of type "tmpfs".
type TmpfsOptions struct {
	// Size sets the size of the tmpfs, in bytes.
	//
	// This will be converted to an operating system specific value
	// depending on the host. For example, on linux, it will be converted to
	// use a 'k', 'm' or 'g' syntax. BSD, though not widely supported with
	// docker, uses a straight byte value.
	//
	// Percentages are not supported.
	SizeBytes int64 `json:",omitempty"`
	// Mode of the tmpfs upon creation
	Mode uint32 `json:",omitempty"`

	// TODO(stevvooe): There are several more tmpfs flags, specified in the
	// daemon, that are accepted. Only the most basic are added for now.
	//
	// From docker/docker/pkg/mount/flags.go:
	//
	// var validFlags = map[string]bool{
	// 	"":          true,
	// 	"size":      true, X
	// 	"mode":      true, X
	// 	"uid":       true,
	// 	"gid":       true,
	// 	"nr_inodes": true,
	// 	"nr_blocks": true,
	// 	"mpol":      true,
	// }
	//
	// Some of these may be straightforward to add, but others, such as
	// uid/gid have implications in a clustered system.
}

// Mount represents a mount (volume).
type Mount struct {
	Type Type `json:",omitempty"`
	// Source specifies the name of the mount. Depending on mount type, this
	// may be a volume name or a host path, or even ignored.
	// Source is not supported for tmpfs (must be an empty value)
	Source      string      `json:",omitempty"`
	Target      string      `json:",omitempty"`
	ReadOnly    bool        `json:",omitempty"`
	Consistency Consistency `json:",omitempty"`

	BindOptions   BindOptions   `json:",omitempty"`
	VolumeOptions VolumeOptions `json:",omitempty"`
	TmpfsOptions  TmpfsOptions  `json:",omitempty"`
}

// HostConfig the non-portable Config structure of a container.
// Here, "non-portable" means "dependent of the host we are running on".
// Portable information *should* appear in Config.
type HostConfig struct {
	// Applicable to all platforms
	Binds           []string      // List of volume bindings for this container
	ContainerIDFile string        // File (path) where the containerId is written
	LogConfig       LogConfig     // Configuration of the logs for this container
	NetworkMode     NetworkMode   // Network mode to use for the container
	PortBindings    PortMap       // Port mapping between the exposed port (container) and the host
	RestartPolicy   RestartPolicy // Restart policy to be used for the container
	AutoRemove      bool          // Automatically remove container when it exits
	VolumeDriver    string        // Name of the volume driver used to mount volumes
	VolumesFrom     []string      // List of volumes to take from other container

	// Applicable to UNIX platforms
	CapAdd          StrSlice          // List of kernel capabilities to add to the container
	CapDrop         StrSlice          // List of kernel capabilities to remove from the container
	Capabilities    []string          `json:"Capabilities"` // List of kernel capabilities to be available for container (this overrides the default set)
	DNS             []string          `json:"Dns"`          // List of DNS server to lookup
	DNSOptions      []string          `json:"DnsOptions"`   // List of DNSOption to look for
	DNSSearch       []string          `json:"DnsSearch"`    // List of DNSSearch to look for
	ExtraHosts      []string          // List of extra hosts
	GroupAdd        []string          // List of additional groups that the container process will run as
	IpcMode         IpcMode           // IPC namespace to use for the container
	Cgroup          CgroupSpec        // Cgroup to use for the container
	Links           []string          // List of links (in the name:alias form)
	OomScoreAdj     int               // Container preference for OOM-killing
	PidMode         PidMode           // PID namespace to use for the container
	Privileged      bool              // Is the container in privileged mode
	PublishAllPorts bool              // Should docker publish all exposed port for the container
	ReadonlyRootfs  bool              // Is the container root filesystem in read-only
	SecurityOpt     []string          // List of string values to customize labels for MLS systems, such as SELinux.
	StorageOpt      map[string]string `json:",omitempty"` // Storage driver options per container.
	Tmpfs           map[string]string `json:",omitempty"` // List of tmpfs (mounts) used for the container
	UTSMode         UTSMode           // UTS namespace to use for the container
	UsernsMode      UsernsMode        // The user namespace to use for the container
	ShmSize         int64             // Total shm memory usage
	Sysctls         map[string]string `json:",omitempty"` // List of Namespaced sysctls used for the container
	Runtime         string            `json:",omitempty"` // Runtime to use with this container

	// Applicable to Windows
	ConsoleSize [2]uint   // Initial console size (height,width)
	Isolation   Isolation // Isolation technology of the container (e.g. default, hyperv)

	// Contains container's resources (cgroups, ulimits)
	Resources

	// Mounts specs used by the container
	Mounts []Mount `json:",omitempty"`

	// MaskedPaths is the list of paths to be masked inside the container (this overrides the default set of paths)
	MaskedPaths []string

	// ReadonlyPaths is the list of paths to be set as read-only inside the container (this overrides the default set of paths)
	ReadonlyPaths []string

	// Run a custom init inside the container, if null, use the daemon's configured settings
	Init bool `json:",omitempty"`
}

// GraphDriverData Information about a container's graph driver.
// swagger:model GraphDriverData
type GraphDriverData struct {

	// data
	// Required: true
	Data map[string]string `json:"Data"`

	// name
	// Required: true
	Name string `json:"Name"`
}

type ContainerJSONBase struct {
	ID              string `json:"Id"`
	Created         string
	Path            string
	Args            []string
	State           ContainerState
	Image           string
	ResolvConfPath  string
	HostnamePath    string
	HostsPath       string
	LogPath         string
	Node            ContainerNode `json:",omitempty"`
	Name            string
	RestartCount    int
	Driver          string
	Platform        string
	MountLabel      string
	ProcessLabel    string
	AppArmorProfile string
	ExecIDs         []string
	HostConfig      HostConfig
	GraphDriver     GraphDriverData
	SizeRw          int64 `json:",omitempty"`
	SizeRootFs      int64 `json:",omitempty"`
}

// ContainerNode stores information about the node that a container
// is running on.  It's only available in Docker Swarm
type ContainerNode struct {
	ID        string
	IPAddress string `json:"IP"`
	Addr      string
	Name      string
	Cpus      int
	Memory    int64
	Labels    map[string]string
}

// Type represents the type of a mount.
type Type string

// Propagation represents the propagation of a mount.
type Propagation string

// MountPoint represents a mount point configuration inside the container.
// This is used for reporting the mountpoints in use by a container.
type MountPoint struct {
	Type        Type   `json:",omitempty"`
	Name        string `json:",omitempty"`
	Source      string
	Destination string
	Driver      string `json:",omitempty"`
	Mode        string
	RW          bool
	Propagation Propagation
}

// HealthConfig holds configuration settings for the HEALTHCHECK feature.
type HealthConfig struct {
	// Test is the test to perform to check that the container is healthy.
	// An empty slice means to inherit the default.
	// The options are:
	// {} : inherit healthcheck
	// {"NONE"} : disable healthcheck
	// {"CMD", args...} : exec arguments directly
	// {"CMD-SHELL", command} : run command with system's default shell
	Test []string `json:",omitempty"`

	// Zero means to inherit. Durations are expressed as integer nanoseconds.
	Interval    int64 `json:",omitempty"` // Interval is the time to wait between checks.
	Timeout     int64 `json:",omitempty"` // Timeout is the time to wait before considering the check to have hung.
	StartPeriod int64 `json:",omitempty"` // The start period for the container to initialize before the retries starts to count down.

	// Retries is the number of consecutive failures needed to consider a container as unhealthy.
	// Zero means inherit.
	Retries int `json:",omitempty"`
}

// Config contains the configuration data about a container.
// It should hold only portable information about the container.
// Here, "portable" means "independent from the host we are running on".
// Non-portable information *should* appear in HostConfig.
// All fields added to this struct must be marked `omitempty` to keep getting
// predictable hashes from the old `v1Compatibility` configuration.
type Config struct {
	Hostname        string              // Hostname
	Domainname      string              // Domainname
	User            string              // User that will run the command(s) inside the container, also support user:group
	AttachStdin     bool                // Attach the standard input, makes possible user interaction
	AttachStdout    bool                // Attach the standard output
	AttachStderr    bool                // Attach the standard error
	ExposedPorts    PortSet             `json:",omitempty"` // List of exposed ports
	Tty             bool                // Attach standard streams to a tty, including stdin if it is not closed.
	OpenStdin       bool                // Open stdin
	StdinOnce       bool                // If true, close stdin after the 1 attached client disconnects.
	Env             []string            // List of environment variable to set in the container
	Cmd             StrSlice            // Command to run when starting the container
	Healthcheck     HealthConfig        `json:",omitempty"` // Healthcheck describes how to check the container is healthy
	ArgsEscaped     bool                `json:",omitempty"` // True if command is already escaped (meaning treat as a command line) (Windows specific).
	Image           string              // Name of the image as it was passed by the operator (e.g. could be symbolic)
	Volumes         map[string]struct{} // List of volumes (mounts) used for the container
	WorkingDir      string              // Current directory (PWD) in the command will be launched
	Entrypoint      StrSlice            // Entrypoint to run when starting the container
	NetworkDisabled bool                `json:",omitempty"` // Is network disabled
	MacAddress      string              `json:",omitempty"` // Mac Address of the container
	OnBuild         []string            // ONBUILD metadata that were defined on the image Dockerfile
	Labels          map[string]string   // List of labels set to this container
	StopSignal      string              `json:",omitempty"` // Signal to stop a container
	StopTimeout     int                 `json:",omitempty"` // Timeout (in seconds) to stop a container
	Shell           StrSlice            `json:",omitempty"` // Shell for shell-form of RUN, CMD, ENTRYPOINT
}

// Address represents an IP address
type Address struct {
	Addr      string
	PrefixLen int
}

// PortBinding represents a binding between a Host IP address and a Host Port
type PortBinding struct {
	// HostIP is the host IP Address
	HostIP string `json:"HostIp"`
	// HostPort is the host port number
	HostPort string
}

// PortMap is a collection of PortBinding indexed by Port
type PortMap map[Port][]PortBinding

// PortSet is a collection of structs indexed by Port
type PortSet map[Port]struct{}

// Port is a string containing port number and protocol in the format "80/tcp"
type Port string

// NetworkSettingsBase holds basic information about networks
type NetworkSettingsBase struct {
	Bridge                 string  // Bridge is the Bridge name the network uses(e.g. `docker0`)
	SandboxID              string  // SandboxID uniquely represents a container's network stack
	HairpinMode            bool    // HairpinMode specifies if hairpin NAT should be enabled on the virtual interface
	LinkLocalIPv6Address   string  // LinkLocalIPv6Address is an IPv6 unicast address using the link-local prefix
	LinkLocalIPv6PrefixLen int     // LinkLocalIPv6PrefixLen is the prefix length of an IPv6 unicast address
	Ports                  PortMap // Ports is a collection of PortBinding indexed by Port
	SandboxKey             string  // SandboxKey identifies the sandbox
	SecondaryIPAddresses   []Address
	SecondaryIPv6Addresses []Address
}

// EndpointIPAMConfig represents IPAM configurations for the endpoint
type EndpointIPAMConfig struct {
	IPv4Address  string   `json:",omitempty"`
	IPv6Address  string   `json:",omitempty"`
	LinkLocalIPs []string `json:",omitempty"`
}

// EndpointSettings stores the network endpoint details
type EndpointSettings struct {
	// Configurations
	IPAMConfig EndpointIPAMConfig
	Links      []string
	Aliases    []string
	// Operational data
	NetworkID           string
	EndpointID          string
	Gateway             string
	IPAddress           string
	IPPrefixLen         int
	IPv6Gateway         string
	GlobalIPv6Address   string
	GlobalIPv6PrefixLen int
	MacAddress          string
	DriverOpts          map[string]string
}

type DefaultNetworkSettings struct {
	EndpointID          string // EndpointID uniquely represents a service endpoint in a Sandbox
	Gateway             string // Gateway holds the gateway address for the network
	GlobalIPv6Address   string // GlobalIPv6Address holds network's global IPv6 address
	GlobalIPv6PrefixLen int    // GlobalIPv6PrefixLen represents mask length of network's global IPv6 address
	IPAddress           string // IPAddress holds the IPv4 address for the network
	IPPrefixLen         int    // IPPrefixLen represents mask length of network's IPv4 address
	IPv6Gateway         string // IPv6Gateway holds gateway address specific for IPv6
	MacAddress          string // MacAddress holds the MAC address for the network
}

// NetworkSettings exposes the network settings in the api
type NetworkSettings struct {
	NetworkSettingsBase
	DefaultNetworkSettings
	Networks map[string]EndpointSettings
}

type ContainerJSON struct {
	ContainerJSONBase
	Mounts          []MountPoint
	Config          Config
	NetworkSettings NetworkSettings
}

func main() {

	/*
	  toFile := bytes.Buffer{}
	  a := []string{"bool", "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16",
	    "uint32", "uint64", "uintptr", "float32", "float64", "interface{}", "string"}
	  for _, v := range a {
	    for _, v2 := range a {
	      toFile.WriteString( fmt.Sprintf("case map[%v]%v:\n", v, v2) )
	      toFile.WriteString( fmt.Sprintf("  keyType = \"%v\"\n", v) )
	      toFile.WriteString( fmt.Sprintf("  keyValue = \"%v\"\n", v2) )
	    }
	  }

	  ioutil.WriteFile("./out.txt", toFile.Bytes(), os.ModePerm)

	  os.Exit(0)
	*/

	var err error
	var content = []byte(`
syntax = "proto3";

option go_package = "github.com/helmutkemper/iotmaker_docker_communication_grpc";

package iotmakerDockerCommunicationGrpc;

`)

	err = ioutil.WriteFile("./out.proto", content, os.ModePerm)
	if err != nil {
		panic(err)
	}

	var a1 ContainerJSON
	test(&a1)
	var a2 Mount
	test(&a2)
	var a3 HealthcheckResult
	test(&a3)
	var a4 WeightDevice
	test(&a4)
	var a5 ThrottleDevice
	test(&a5)
	var a6 DeviceMapping
	test(&a6)
	var a7 DeviceRequest
	test(&a7)
	var a8 Ulimit
	test(&a8)
	var a9 MountPoint
	test(&a9)
	var b1 Address
	test(&b1)
	var b2 EndpointSettings
	test(&b2)
	var b3 PortBinding
	test(&b3)

}

func test(i interface{}) {
	var err error
	var file *os.File
	var buffer bytes.Buffer

	elementName := reflect.ValueOf(i).Elem().Type().Name()
	element := reflect.ValueOf(i).Elem()

	buffer.WriteString(fmt.Sprintf("message %v {\n", elementName))

	for i := 0; i < element.NumField(); i += 1 {
		field := element.Field(i)
		nameOfField := element.Type().Field(i).Name

		err = ToScalarValue(&buffer, field, nameOfField, i)
		if err != nil {
			panic(err)
		}
	}

	buffer.WriteString("}\n")
	buffer.WriteString("\n")

	file, err = os.OpenFile("./out.proto", os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	_, err = file.Write(buffer.Bytes())
	if err != nil {
		panic(err)
	}
}

func ToScalarValue(
	buffer *bytes.Buffer,
	element reflect.Value,
	nameOfField string,
	i int,
) (
	err error,
) {

	nameOfField = strings.Replace(nameOfField, "main.", "", -1)
	var elementTypeText = element.Type().String()
	elementTypeText = strings.Replace(elementTypeText, "main.", "", -1)
	switch element.Type().Kind() {
	case reflect.Invalid:
		err = errors.New("ToScalarValue() function found an invalid value")
	case reflect.Bool:
		buffer.WriteString("  " + elementTypeText + " " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case reflect.Uintptr:
		buffer.WriteString("  " + removePtrFromString(elementTypeText) + " " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case reflect.Complex64:
		err = errors.New("ToScalarValue() function has't Complex64 code")
	case reflect.Complex128:
		err = errors.New("ToScalarValue() function has't Complex128 code")
	case reflect.Array:
		//err = errors.New("ToScalarValue() function has't array code. >"+nameOfField)
	case reflect.Chan:
		break
	case reflect.Func:
		break
	case reflect.Interface:
		buffer.WriteString("  interface " + elementTypeText + " " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case reflect.Map:
		mapConverter(buffer, element, nameOfField, i)
	case reflect.Struct:
		err = ToStructType(element)
		buffer.WriteString("  " + elementTypeText + " " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case reflect.UnsafePointer:
		break
	case reflect.String:
		buffer.WriteString("  string " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case reflect.Uint:
		buffer.WriteString("  uint64 " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case reflect.Uint8:
		buffer.WriteString("  uint32 " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case reflect.Uint16:
		buffer.WriteString("  uint32 " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case reflect.Uint32:
		buffer.WriteString("  uint32 " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case reflect.Uint64:
		buffer.WriteString("  uint64 " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case reflect.Int:
		buffer.WriteString("  int64 " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case reflect.Int8:
		buffer.WriteString("  int8 " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case reflect.Int16:
		buffer.WriteString("  int16 " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case reflect.Int32:
		buffer.WriteString("  int32 " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case reflect.Int64:
		buffer.WriteString("  int64 " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case reflect.Float32:
		buffer.WriteString("  float " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case reflect.Float64:
		buffer.WriteString("  double " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case reflect.Slice:
		buffer.WriteString("  repeated " + removeArrFromString(elementTypeText) + " " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case reflect.Ptr:
		buffer.WriteString("  " + removePtrFromString(elementTypeText) + " " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	}

	return
}

func ToStructType(
	//buffer *bytes.Buffer,
	element reflect.Value,
) (
	err error,
) {

	var file *os.File
	var buffer bytes.Buffer

	//var elementTypeText = element.Type().String()
	var nameOfStruct = element.Type().Name()
	_ = nameOfStruct
	t := element
	_ = t

	buffer.WriteString(fmt.Sprintf("message %v {\n", nameOfStruct))

	for i := 0; i < element.NumField(); i += 1 {
		nameOfField := element.Type().Field(i).Name
		field := element.Field(i)
		err = ToScalarValue(&buffer, field, nameOfField, i)
		if err != nil {
			return
		}
	}

	buffer.WriteString("}\n")
	buffer.WriteString("\n")

	file, err = os.OpenFile("./out.proto", os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	_, err = file.Write(buffer.Bytes())
	if err != nil {
		panic(err)
	}
	return
}

func removeArrFromString(t string) (convertedType string) {
	return strings.Replace(t, "[]", "", 1)
}

func removePtrFromString(t string) (convertedType string) {
	return strings.Replace(t, "*", "", 1)
}

func mapConverter(buffer *bytes.Buffer, element reflect.Value, nameOfField string, i int) {
	var keyType, keyValue string

	switch element.Interface().(type) {
	case map[bool]bool:
		keyType = "bool"
		keyValue = "bool"
	case map[bool]int:
		keyType = "bool"
		keyValue = "int"
	case map[bool]int8:
		keyType = "bool"
		keyValue = "int8"
	case map[bool]int16:
		keyType = "bool"
		keyValue = "int16"
	case map[bool]int32:
		keyType = "bool"
		keyValue = "int32"
	case map[bool]int64:
		keyType = "bool"
		keyValue = "int64"
	case map[bool]uint:
		keyType = "bool"
		keyValue = "uint"
	case map[bool]uint8:
		keyType = "bool"
		keyValue = "uint8"
	case map[bool]uint16:
		keyType = "bool"
		keyValue = "uint16"
	case map[bool]uint32:
		keyType = "bool"
		keyValue = "uint32"
	case map[bool]uint64:
		keyType = "bool"
		keyValue = "uint64"
	case map[bool]uintptr:
		keyType = "bool"
		keyValue = "uintptr"
	case map[bool]float32:
		keyType = "bool"
		keyValue = "float32"
	case map[bool]float64:
		keyType = "bool"
		keyValue = "float64"
	case map[bool]interface{}:
		keyType = "bool"
		keyValue = "interface{}"
	case map[bool]string:
		keyType = "bool"
		keyValue = "string"
	case map[int]bool:
		keyType = "int"
		keyValue = "bool"
	case map[int]int:
		keyType = "int"
		keyValue = "int"
	case map[int]int8:
		keyType = "int"
		keyValue = "int8"
	case map[int]int16:
		keyType = "int"
		keyValue = "int16"
	case map[int]int32:
		keyType = "int"
		keyValue = "int32"
	case map[int]int64:
		keyType = "int"
		keyValue = "int64"
	case map[int]uint:
		keyType = "int"
		keyValue = "uint"
	case map[int]uint8:
		keyType = "int"
		keyValue = "uint8"
	case map[int]uint16:
		keyType = "int"
		keyValue = "uint16"
	case map[int]uint32:
		keyType = "int"
		keyValue = "uint32"
	case map[int]uint64:
		keyType = "int"
		keyValue = "uint64"
	case map[int]uintptr:
		keyType = "int"
		keyValue = "uintptr"
	case map[int]float32:
		keyType = "int"
		keyValue = "float32"
	case map[int]float64:
		keyType = "int"
		keyValue = "float64"
	case map[int]interface{}:
		keyType = "int"
		keyValue = "interface{}"
	case map[int]string:
		keyType = "int"
		keyValue = "string"
	case map[int8]bool:
		keyType = "int8"
		keyValue = "bool"
	case map[int8]int:
		keyType = "int8"
		keyValue = "int"
	case map[int8]int8:
		keyType = "int8"
		keyValue = "int8"
	case map[int8]int16:
		keyType = "int8"
		keyValue = "int16"
	case map[int8]int32:
		keyType = "int8"
		keyValue = "int32"
	case map[int8]int64:
		keyType = "int8"
		keyValue = "int64"
	case map[int8]uint:
		keyType = "int8"
		keyValue = "uint"
	case map[int8]uint8:
		keyType = "int8"
		keyValue = "uint8"
	case map[int8]uint16:
		keyType = "int8"
		keyValue = "uint16"
	case map[int8]uint32:
		keyType = "int8"
		keyValue = "uint32"
	case map[int8]uint64:
		keyType = "int8"
		keyValue = "uint64"
	case map[int8]uintptr:
		keyType = "int8"
		keyValue = "uintptr"
	case map[int8]float32:
		keyType = "int8"
		keyValue = "float32"
	case map[int8]float64:
		keyType = "int8"
		keyValue = "float64"
	case map[int8]interface{}:
		keyType = "int8"
		keyValue = "interface{}"
	case map[int8]string:
		keyType = "int8"
		keyValue = "string"
	case map[int16]bool:
		keyType = "int16"
		keyValue = "bool"
	case map[int16]int:
		keyType = "int16"
		keyValue = "int"
	case map[int16]int8:
		keyType = "int16"
		keyValue = "int8"
	case map[int16]int16:
		keyType = "int16"
		keyValue = "int16"
	case map[int16]int32:
		keyType = "int16"
		keyValue = "int32"
	case map[int16]int64:
		keyType = "int16"
		keyValue = "int64"
	case map[int16]uint:
		keyType = "int16"
		keyValue = "uint"
	case map[int16]uint8:
		keyType = "int16"
		keyValue = "uint8"
	case map[int16]uint16:
		keyType = "int16"
		keyValue = "uint16"
	case map[int16]uint32:
		keyType = "int16"
		keyValue = "uint32"
	case map[int16]uint64:
		keyType = "int16"
		keyValue = "uint64"
	case map[int16]uintptr:
		keyType = "int16"
		keyValue = "uintptr"
	case map[int16]float32:
		keyType = "int16"
		keyValue = "float32"
	case map[int16]float64:
		keyType = "int16"
		keyValue = "float64"
	case map[int16]interface{}:
		keyType = "int16"
		keyValue = "interface{}"
	case map[int16]string:
		keyType = "int16"
		keyValue = "string"
	case map[int32]bool:
		keyType = "int32"
		keyValue = "bool"
	case map[int32]int:
		keyType = "int32"
		keyValue = "int"
	case map[int32]int8:
		keyType = "int32"
		keyValue = "int8"
	case map[int32]int16:
		keyType = "int32"
		keyValue = "int16"
	case map[int32]int32:
		keyType = "int32"
		keyValue = "int32"
	case map[int32]int64:
		keyType = "int32"
		keyValue = "int64"
	case map[int32]uint:
		keyType = "int32"
		keyValue = "uint"
	case map[int32]uint8:
		keyType = "int32"
		keyValue = "uint8"
	case map[int32]uint16:
		keyType = "int32"
		keyValue = "uint16"
	case map[int32]uint32:
		keyType = "int32"
		keyValue = "uint32"
	case map[int32]uint64:
		keyType = "int32"
		keyValue = "uint64"
	case map[int32]uintptr:
		keyType = "int32"
		keyValue = "uintptr"
	case map[int32]float32:
		keyType = "int32"
		keyValue = "float32"
	case map[int32]float64:
		keyType = "int32"
		keyValue = "float64"
	case map[int32]interface{}:
		keyType = "int32"
		keyValue = "interface{}"
	case map[int32]string:
		keyType = "int32"
		keyValue = "string"
	case map[int64]bool:
		keyType = "int64"
		keyValue = "bool"
	case map[int64]int:
		keyType = "int64"
		keyValue = "int"
	case map[int64]int8:
		keyType = "int64"
		keyValue = "int8"
	case map[int64]int16:
		keyType = "int64"
		keyValue = "int16"
	case map[int64]int32:
		keyType = "int64"
		keyValue = "int32"
	case map[int64]int64:
		keyType = "int64"
		keyValue = "int64"
	case map[int64]uint:
		keyType = "int64"
		keyValue = "uint"
	case map[int64]uint8:
		keyType = "int64"
		keyValue = "uint8"
	case map[int64]uint16:
		keyType = "int64"
		keyValue = "uint16"
	case map[int64]uint32:
		keyType = "int64"
		keyValue = "uint32"
	case map[int64]uint64:
		keyType = "int64"
		keyValue = "uint64"
	case map[int64]uintptr:
		keyType = "int64"
		keyValue = "uintptr"
	case map[int64]float32:
		keyType = "int64"
		keyValue = "float32"
	case map[int64]float64:
		keyType = "int64"
		keyValue = "float64"
	case map[int64]interface{}:
		keyType = "int64"
		keyValue = "interface{}"
	case map[int64]string:
		keyType = "int64"
		keyValue = "string"
	case map[uint]bool:
		keyType = "uint"
		keyValue = "bool"
	case map[uint]int:
		keyType = "uint"
		keyValue = "int"
	case map[uint]int8:
		keyType = "uint"
		keyValue = "int8"
	case map[uint]int16:
		keyType = "uint"
		keyValue = "int16"
	case map[uint]int32:
		keyType = "uint"
		keyValue = "int32"
	case map[uint]int64:
		keyType = "uint"
		keyValue = "int64"
	case map[uint]uint:
		keyType = "uint"
		keyValue = "uint"
	case map[uint]uint8:
		keyType = "uint"
		keyValue = "uint8"
	case map[uint]uint16:
		keyType = "uint"
		keyValue = "uint16"
	case map[uint]uint32:
		keyType = "uint"
		keyValue = "uint32"
	case map[uint]uint64:
		keyType = "uint"
		keyValue = "uint64"
	case map[uint]uintptr:
		keyType = "uint"
		keyValue = "uintptr"
	case map[uint]float32:
		keyType = "uint"
		keyValue = "float32"
	case map[uint]float64:
		keyType = "uint"
		keyValue = "float64"
	case map[uint]interface{}:
		keyType = "uint"
		keyValue = "interface{}"
	case map[uint]string:
		keyType = "uint"
		keyValue = "string"
	case map[uint8]bool:
		keyType = "uint8"
		keyValue = "bool"
	case map[uint8]int:
		keyType = "uint8"
		keyValue = "int"
	case map[uint8]int8:
		keyType = "uint8"
		keyValue = "int8"
	case map[uint8]int16:
		keyType = "uint8"
		keyValue = "int16"
	case map[uint8]int32:
		keyType = "uint8"
		keyValue = "int32"
	case map[uint8]int64:
		keyType = "uint8"
		keyValue = "int64"
	case map[uint8]uint:
		keyType = "uint8"
		keyValue = "uint"
	case map[uint8]uint8:
		keyType = "uint8"
		keyValue = "uint8"
	case map[uint8]uint16:
		keyType = "uint8"
		keyValue = "uint16"
	case map[uint8]uint32:
		keyType = "uint8"
		keyValue = "uint32"
	case map[uint8]uint64:
		keyType = "uint8"
		keyValue = "uint64"
	case map[uint8]uintptr:
		keyType = "uint8"
		keyValue = "uintptr"
	case map[uint8]float32:
		keyType = "uint8"
		keyValue = "float32"
	case map[uint8]float64:
		keyType = "uint8"
		keyValue = "float64"
	case map[uint8]interface{}:
		keyType = "uint8"
		keyValue = "interface{}"
	case map[uint8]string:
		keyType = "uint8"
		keyValue = "string"
	case map[uint16]bool:
		keyType = "uint16"
		keyValue = "bool"
	case map[uint16]int:
		keyType = "uint16"
		keyValue = "int"
	case map[uint16]int8:
		keyType = "uint16"
		keyValue = "int8"
	case map[uint16]int16:
		keyType = "uint16"
		keyValue = "int16"
	case map[uint16]int32:
		keyType = "uint16"
		keyValue = "int32"
	case map[uint16]int64:
		keyType = "uint16"
		keyValue = "int64"
	case map[uint16]uint:
		keyType = "uint16"
		keyValue = "uint"
	case map[uint16]uint8:
		keyType = "uint16"
		keyValue = "uint8"
	case map[uint16]uint16:
		keyType = "uint16"
		keyValue = "uint16"
	case map[uint16]uint32:
		keyType = "uint16"
		keyValue = "uint32"
	case map[uint16]uint64:
		keyType = "uint16"
		keyValue = "uint64"
	case map[uint16]uintptr:
		keyType = "uint16"
		keyValue = "uintptr"
	case map[uint16]float32:
		keyType = "uint16"
		keyValue = "float32"
	case map[uint16]float64:
		keyType = "uint16"
		keyValue = "float64"
	case map[uint16]interface{}:
		keyType = "uint16"
		keyValue = "interface{}"
	case map[uint16]string:
		keyType = "uint16"
		keyValue = "string"
	case map[uint32]bool:
		keyType = "uint32"
		keyValue = "bool"
	case map[uint32]int:
		keyType = "uint32"
		keyValue = "int"
	case map[uint32]int8:
		keyType = "uint32"
		keyValue = "int8"
	case map[uint32]int16:
		keyType = "uint32"
		keyValue = "int16"
	case map[uint32]int32:
		keyType = "uint32"
		keyValue = "int32"
	case map[uint32]int64:
		keyType = "uint32"
		keyValue = "int64"
	case map[uint32]uint:
		keyType = "uint32"
		keyValue = "uint"
	case map[uint32]uint8:
		keyType = "uint32"
		keyValue = "uint8"
	case map[uint32]uint16:
		keyType = "uint32"
		keyValue = "uint16"
	case map[uint32]uint32:
		keyType = "uint32"
		keyValue = "uint32"
	case map[uint32]uint64:
		keyType = "uint32"
		keyValue = "uint64"
	case map[uint32]uintptr:
		keyType = "uint32"
		keyValue = "uintptr"
	case map[uint32]float32:
		keyType = "uint32"
		keyValue = "float32"
	case map[uint32]float64:
		keyType = "uint32"
		keyValue = "float64"
	case map[uint32]interface{}:
		keyType = "uint32"
		keyValue = "interface{}"
	case map[uint32]string:
		keyType = "uint32"
		keyValue = "string"
	case map[uint64]bool:
		keyType = "uint64"
		keyValue = "bool"
	case map[uint64]int:
		keyType = "uint64"
		keyValue = "int"
	case map[uint64]int8:
		keyType = "uint64"
		keyValue = "int8"
	case map[uint64]int16:
		keyType = "uint64"
		keyValue = "int16"
	case map[uint64]int32:
		keyType = "uint64"
		keyValue = "int32"
	case map[uint64]int64:
		keyType = "uint64"
		keyValue = "int64"
	case map[uint64]uint:
		keyType = "uint64"
		keyValue = "uint"
	case map[uint64]uint8:
		keyType = "uint64"
		keyValue = "uint8"
	case map[uint64]uint16:
		keyType = "uint64"
		keyValue = "uint16"
	case map[uint64]uint32:
		keyType = "uint64"
		keyValue = "uint32"
	case map[uint64]uint64:
		keyType = "uint64"
		keyValue = "uint64"
	case map[uint64]uintptr:
		keyType = "uint64"
		keyValue = "uintptr"
	case map[uint64]float32:
		keyType = "uint64"
		keyValue = "float32"
	case map[uint64]float64:
		keyType = "uint64"
		keyValue = "float64"
	case map[uint64]interface{}:
		keyType = "uint64"
		keyValue = "interface{}"
	case map[uint64]string:
		keyType = "uint64"
		keyValue = "string"
	case map[uintptr]bool:
		keyType = "uintptr"
		keyValue = "bool"
	case map[uintptr]int:
		keyType = "uintptr"
		keyValue = "int"
	case map[uintptr]int8:
		keyType = "uintptr"
		keyValue = "int8"
	case map[uintptr]int16:
		keyType = "uintptr"
		keyValue = "int16"
	case map[uintptr]int32:
		keyType = "uintptr"
		keyValue = "int32"
	case map[uintptr]int64:
		keyType = "uintptr"
		keyValue = "int64"
	case map[uintptr]uint:
		keyType = "uintptr"
		keyValue = "uint"
	case map[uintptr]uint8:
		keyType = "uintptr"
		keyValue = "uint8"
	case map[uintptr]uint16:
		keyType = "uintptr"
		keyValue = "uint16"
	case map[uintptr]uint32:
		keyType = "uintptr"
		keyValue = "uint32"
	case map[uintptr]uint64:
		keyType = "uintptr"
		keyValue = "uint64"
	case map[uintptr]uintptr:
		keyType = "uintptr"
		keyValue = "uintptr"
	case map[uintptr]float32:
		keyType = "uintptr"
		keyValue = "float32"
	case map[uintptr]float64:
		keyType = "uintptr"
		keyValue = "float64"
	case map[uintptr]interface{}:
		keyType = "uintptr"
		keyValue = "interface{}"
	case map[uintptr]string:
		keyType = "uintptr"
		keyValue = "string"
	case map[float32]bool:
		keyType = "float32"
		keyValue = "bool"
	case map[float32]int:
		keyType = "float32"
		keyValue = "int"
	case map[float32]int8:
		keyType = "float32"
		keyValue = "int8"
	case map[float32]int16:
		keyType = "float32"
		keyValue = "int16"
	case map[float32]int32:
		keyType = "float32"
		keyValue = "int32"
	case map[float32]int64:
		keyType = "float32"
		keyValue = "int64"
	case map[float32]uint:
		keyType = "float32"
		keyValue = "uint"
	case map[float32]uint8:
		keyType = "float32"
		keyValue = "uint8"
	case map[float32]uint16:
		keyType = "float32"
		keyValue = "uint16"
	case map[float32]uint32:
		keyType = "float32"
		keyValue = "uint32"
	case map[float32]uint64:
		keyType = "float32"
		keyValue = "uint64"
	case map[float32]uintptr:
		keyType = "float32"
		keyValue = "uintptr"
	case map[float32]float32:
		keyType = "float32"
		keyValue = "float32"
	case map[float32]float64:
		keyType = "float32"
		keyValue = "float64"
	case map[float32]interface{}:
		keyType = "float32"
		keyValue = "interface{}"
	case map[float32]string:
		keyType = "float32"
		keyValue = "string"
	case map[float64]bool:
		keyType = "float64"
		keyValue = "bool"
	case map[float64]int:
		keyType = "float64"
		keyValue = "int"
	case map[float64]int8:
		keyType = "float64"
		keyValue = "int8"
	case map[float64]int16:
		keyType = "float64"
		keyValue = "int16"
	case map[float64]int32:
		keyType = "float64"
		keyValue = "int32"
	case map[float64]int64:
		keyType = "float64"
		keyValue = "int64"
	case map[float64]uint:
		keyType = "float64"
		keyValue = "uint"
	case map[float64]uint8:
		keyType = "float64"
		keyValue = "uint8"
	case map[float64]uint16:
		keyType = "float64"
		keyValue = "uint16"
	case map[float64]uint32:
		keyType = "float64"
		keyValue = "uint32"
	case map[float64]uint64:
		keyType = "float64"
		keyValue = "uint64"
	case map[float64]uintptr:
		keyType = "float64"
		keyValue = "uintptr"
	case map[float64]float32:
		keyType = "float64"
		keyValue = "float32"
	case map[float64]float64:
		keyType = "float64"
		keyValue = "float64"
	case map[float64]interface{}:
		keyType = "float64"
		keyValue = "interface{}"
	case map[float64]string:
		keyType = "float64"
		keyValue = "string"
	case map[interface{}]bool:
		keyType = "interface{}"
		keyValue = "bool"
	case map[interface{}]int:
		keyType = "interface{}"
		keyValue = "int"
	case map[interface{}]int8:
		keyType = "interface{}"
		keyValue = "int8"
	case map[interface{}]int16:
		keyType = "interface{}"
		keyValue = "int16"
	case map[interface{}]int32:
		keyType = "interface{}"
		keyValue = "int32"
	case map[interface{}]int64:
		keyType = "interface{}"
		keyValue = "int64"
	case map[interface{}]uint:
		keyType = "interface{}"
		keyValue = "uint"
	case map[interface{}]uint8:
		keyType = "interface{}"
		keyValue = "uint8"
	case map[interface{}]uint16:
		keyType = "interface{}"
		keyValue = "uint16"
	case map[interface{}]uint32:
		keyType = "interface{}"
		keyValue = "uint32"
	case map[interface{}]uint64:
		keyType = "interface{}"
		keyValue = "uint64"
	case map[interface{}]uintptr:
		keyType = "interface{}"
		keyValue = "uintptr"
	case map[interface{}]float32:
		keyType = "interface{}"
		keyValue = "float32"
	case map[interface{}]float64:
		keyType = "interface{}"
		keyValue = "float64"
	case map[interface{}]interface{}:
		keyType = "interface{}"
		keyValue = "interface{}"
	case map[interface{}]string:
		keyType = "interface{}"
		keyValue = "string"
	case map[string]bool:
		keyType = "string"
		keyValue = "bool"
	case map[string]int:
		keyType = "string"
		keyValue = "int"
	case map[string]int8:
		keyType = "string"
		keyValue = "int8"
	case map[string]int16:
		keyType = "string"
		keyValue = "int16"
	case map[string]int32:
		keyType = "string"
		keyValue = "int32"
	case map[string]int64:
		keyType = "string"
		keyValue = "int64"
	case map[string]uint:
		keyType = "string"
		keyValue = "uint"
	case map[string]uint8:
		keyType = "string"
		keyValue = "uint8"
	case map[string]uint16:
		keyType = "string"
		keyValue = "uint16"
	case map[string]uint32:
		keyType = "string"
		keyValue = "uint32"
	case map[string]uint64:
		keyType = "string"
		keyValue = "uint64"
	case map[string]uintptr:
		keyType = "string"
		keyValue = "uintptr"
	case map[string]float32:
		keyType = "string"
		keyValue = "float32"
	case map[string]float64:
		keyType = "string"
		keyValue = "float64"
	case map[string]interface{}:
		keyType = "string"
		keyValue = "interface{}"
	case map[string]string:
		keyType = "string"
		keyValue = "string"
	}

	_ = keyType
	_ = keyValue

	buffer.WriteString("  map<" + keyType + ", " + keyValue + "> " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")

}

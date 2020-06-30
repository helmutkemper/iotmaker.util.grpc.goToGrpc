package server

type DeviceMapping struct {
	PathOnHost        string
	PathInContainer   string
	CgroupPermissions string
}

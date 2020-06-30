package server

type DeviceRequest struct {
	Driver       string
	Count        int
	DeviceIDs    []string
	Capabilities [][]string
	Options      map[string]string
}

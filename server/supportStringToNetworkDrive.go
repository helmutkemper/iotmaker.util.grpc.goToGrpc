package server

import iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.0"

func SupportStringToNetworkDrive(str string) (drive iotmakerdocker.NetworkDrive) {
	switch str {
	case "bridge":
		drive = iotmakerdocker.KNetworkDriveBridge
	case "host":
		drive = iotmakerdocker.KNetworkDriveHost
	case "overlay":
		drive = iotmakerdocker.KNetworkDriveOverlay
	case "macvlan":
		drive = iotmakerdocker.KNetworkDriveMacVLan
	case "none":
		drive = iotmakerdocker.KNetworkDriveNone
	default:
		drive = iotmakerdocker.KNetworkDriveNone
	}

	return
}

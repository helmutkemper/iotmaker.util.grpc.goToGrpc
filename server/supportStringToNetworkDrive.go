package server

import (
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
)

func SupportStringToNetworkDrive(str string) (drive iotmakerDocker.NetworkDrive) {
	switch str {
	case "bridge":
		drive = iotmakerDocker.KNetworkDriveBridge
	case "host":
		drive = iotmakerDocker.KNetworkDriveHost
	case "overlay":
		drive = iotmakerDocker.KNetworkDriveOverlay
	case "macvlan":
		drive = iotmakerDocker.KNetworkDriveMacVLan
	case "none":
		drive = iotmakerDocker.KNetworkDriveNone
	default:
		drive = iotmakerDocker.KNetworkDriveNone
	}

	return
}

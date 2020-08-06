package server

import (
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
	"time"
)

type BuildOrPullLog struct {
	Status iotmakerDocker.ContainerPullStatusSendToChannel
	Log    string
	Start  time.Time
}

package server

import (
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.0"
	"time"
)

type BuildOrPullLog struct {
	Status iotmakerdocker.ContainerPullStatusSendToChannel
	Log    string
	Start  time.Time
}

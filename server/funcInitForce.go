package server

import iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.0"

func (el *GRpcServer) InitForce() (err error) {
	el.dockerSystem = iotmakerdocker.DockerSystem{}
	err = el.dockerSystem.Init()
	return
}

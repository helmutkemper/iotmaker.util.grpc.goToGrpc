package server

import iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.0"

func (el *GRpcServer) Init() (err error) {
	if el.init == true {
		return
	}

	el.dockerSystem = iotmakerdocker.DockerSystem{}
	err = el.dockerSystem.Init()
	return
}

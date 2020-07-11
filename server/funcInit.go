package server

import iotmakerDocker "github.com/helmutkemper/iotmaker.docker"

func (el *GRpcServer) Init() (err error) {
	if el.init == true {
		return
	}

	el.dockerSystem = iotmakerDocker.DockerSystem{}
	err = el.dockerSystem.Init()
	return
}

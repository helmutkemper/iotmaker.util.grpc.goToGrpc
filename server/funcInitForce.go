package server

import iotmakerDocker "github.com/helmutkemper/iotmaker.docker"

func (el *GRpcServer) InitForce() (err error) {
	el.dockerSystem = &iotmakerDocker.DockerSystem{}
	err = el.dockerSystem.Init()
	return
}

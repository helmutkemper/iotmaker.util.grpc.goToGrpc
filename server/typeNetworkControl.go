package server

import iotmakerDocker "github.com/helmutkemper/iotmaker.docker"

type NetworkControl struct {
	Generator *iotmakerDocker.NextNetworkAutoConfiguration
	Name      string
	Drive     string
	Scope     string
	Subnet    string
	Gateway   string
}

var networkControl map[string]NetworkControl

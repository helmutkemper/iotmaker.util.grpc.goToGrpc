package server

import iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.0"

type NetworkControl struct {
	Generator *iotmakerdocker.NextNetworkAutoConfiguration
	ID        string
	Name      string
	Drive     string
	Scope     string
	Subnet    string
	Gateway   string
}

var networkControl map[string]NetworkControl

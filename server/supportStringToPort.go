package server

import (
	"github.com/docker/go-connections/nat"
	"strings"
)

func SupportStringToPort(
	in string,
) (
	err error,
	port nat.Port,
) {

	arr := strings.Split(in, "/")
	portString := arr[0]
	proto := "tcp"
	if len(arr) > 1 {
		proto = arr[1]
	}
	port, err = nat.NewPort(proto, portString)
	return
}

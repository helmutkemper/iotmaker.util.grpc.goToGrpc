package server

import (
	"github.com/docker/go-connections/nat"
	"strings"
)

func SupportStringToPort(
	in string,
) (
	port nat.Port,
	err error,
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

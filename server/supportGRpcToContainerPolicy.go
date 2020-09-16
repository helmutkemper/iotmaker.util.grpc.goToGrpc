package server

import (
	"errors"
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.0"
)

func SupportGRpcToContainerPolicy(
	value string,
) (
	policy iotmakerdocker.RestartPolicy,
	err error,
) {

	switch value {
	case "no":
		return iotmakerdocker.KRestartPolicyNo, nil

	case "on-failure":
		return iotmakerdocker.KRestartPolicyOnFailure, nil

	case "always":
		return iotmakerdocker.KRestartPolicyAlways, nil

	case "unless-stopped":
		return iotmakerdocker.KRestartPolicyUnlessStopped, nil
	}

	return iotmakerdocker.KRestartPolicyNo, errors.New(value + " not implemented")
}

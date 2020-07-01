package server

import (
	"errors"
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
)

func SupportGRpcToContainerPolicy(value string) (err error, policy iotmakerDocker.RestartPolicy) {
	switch value {
	case "no":
		return nil, iotmakerDocker.KRestartPolicyNo

	case "on-failure":
		return nil, iotmakerDocker.KRestartPolicyOnFailure

	case "always":
		return nil, iotmakerDocker.KRestartPolicyAlways

	case "unless-stopped":
		return nil, iotmakerDocker.KRestartPolicyUnlessStopped
	}

	return errors.New(value + " not implemented"), iotmakerDocker.KRestartPolicyNo
}

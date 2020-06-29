package iotmaker_util_grpc_goToGrpc

type HealthcheckResult struct {
	Start    Time
	End      Time
	ExitCode int
	Output   string
}

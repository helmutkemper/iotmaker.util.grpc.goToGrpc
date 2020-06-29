package iotmaker_util_grpc_goToGrpc

type HealthcheckResult struct {
	Start    int64
	End      int64
	ExitCode int
	Output   string
}

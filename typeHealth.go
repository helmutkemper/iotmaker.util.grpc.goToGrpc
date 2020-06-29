package iotmaker_util_grpc_goToGrpc

type Health struct {
	Status        string
	FailingStreak int
	Log           []HealthcheckResult
}

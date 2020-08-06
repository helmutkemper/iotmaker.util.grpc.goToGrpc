package server

type HealthCheckResult struct {
	Start    int64
	End      int64
	ExitCode int
	Output   string
}

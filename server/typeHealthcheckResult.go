package server

type HealthcheckResult struct {
	Start    int64
	End      int64
	ExitCode int
	Output   string
}

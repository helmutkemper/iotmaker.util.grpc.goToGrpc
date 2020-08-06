package server

type Health struct {
	Status        string
	FailingStreak int
	Log           []HealthCheckResult
}

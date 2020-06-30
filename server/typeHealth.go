package server

type Health struct {
	Status        string
	FailingStreak int
	Log           []HealthcheckResult
}

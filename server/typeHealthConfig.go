package server

type HealthConfig struct {
	Test        []string
	Interval    int64
	Timeout     int64
	StartPeriod int64
	Retries     int
}

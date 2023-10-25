package healthdom

import "context"

type HealthUseCaseInterface interface {
	Ping() string
	HealthCheck(ctx context.Context) DatabaseHealth
}

type HealthRepoInterface interface {
	PingPostgres(ctx context.Context) error
}

type Ping struct {
	Message string `json:"message"`
}

type DatabaseHealth struct {
	PostgresReachable bool `json:"pg_reachable"`
}

package healthuc

import (
	"context"

	healthdom "github.com/mrpandey/gobp/src/internal/core/domain/health"
	"github.com/mrpandey/gobp/src/util"
)

// healthUseCase implements HealthUseCaseInterface.
type healthUseCase struct {
	logger *util.StandardLogger
	repos  healthdom.HealthRepoInterface
}

func NewHealthUseCase(logger *util.StandardLogger, repos healthdom.HealthRepoInterface) *healthUseCase {
	return &healthUseCase{
		logger: logger,
		repos:  repos,
	}
}

func (uc *healthUseCase) Ping() string {
	return "pong"
}

func (uc *healthUseCase) HealthCheck(ctx context.Context) healthdom.DatabaseHealth {
	pgErr := uc.repos.PingPostgres(ctx)
	return healthdom.DatabaseHealth{
		PostgresReachable: pgErr == nil,
	}
}

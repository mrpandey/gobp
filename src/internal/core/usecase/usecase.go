package usecase

import (
	authdom "github.com/mrpandey/gobp/src/internal/core/domain/auth"
	fdom "github.com/mrpandey/gobp/src/internal/core/domain/furniture"
	healthdom "github.com/mrpandey/gobp/src/internal/core/domain/health"

	authuc "github.com/mrpandey/gobp/src/internal/core/usecase/auth"
	fuc "github.com/mrpandey/gobp/src/internal/core/usecase/furniture"
	healthuc "github.com/mrpandey/gobp/src/internal/core/usecase/health"

	"github.com/mrpandey/gobp/src/internal/repo"
	"github.com/mrpandey/gobp/src/util"
)

type UseCases struct {
	Auth      authdom.AuthUseCaseInterface
	Health    healthdom.HealthUseCaseInterface
	Furniture fdom.FurnitureUseCaseInterface
}

func InitUseCases(
	cfg *util.Config,
	logger *util.StandardLogger,
	repos *repo.Repo,
) *UseCases {
	return &UseCases{
		Auth:      authuc.NewAuthUseCase(cfg, logger, repos.TxnManager, repos.Auth),
		Health:    healthuc.NewHealthUseCase(logger, repos.Health),
		Furniture: fuc.NewFurnitureUseCase(logger, repos.TxnManager, repos.Furniture),
	}
}

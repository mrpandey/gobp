package grpcm

import (
	authdom "github.com/mrpandey/gobp/src/internal/core/domain/auth"
	"github.com/mrpandey/gobp/src/util"
)

type Middleware struct {
	cfg         *util.Config
	logger      *util.StandardLogger
	authUsecase authdom.AuthUseCaseInterface
}

func NewMiddleware(
	config *util.Config,
	logger *util.StandardLogger,
	authUsecase authdom.AuthUseCaseInterface,
) *Middleware {
	return &Middleware{
		cfg:         config,
		logger:      logger,
		authUsecase: authUsecase,
	}
}

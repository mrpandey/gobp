package resth

import (
	"github.com/mrpandey/gobp/src/internal/core/usecase"
	"github.com/mrpandey/gobp/src/util"
)

type RestServer struct {
	logger    *util.StandardLogger
	validator *util.Validator
	useCases  *usecase.UseCases
}

func NewRestServer(
	logger *util.StandardLogger,
	validator *util.Validator,
	uc *usecase.UseCases,
) *RestServer {
	return &RestServer{
		logger:    logger,
		validator: validator,
		useCases:  uc,
	}
}

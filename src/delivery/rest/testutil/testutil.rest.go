package testrest

import (
	"testing"

	resth "github.com/mrpandey/gobp/src/delivery/rest/handlers"
	"github.com/mrpandey/gobp/src/internal/core/usecase"
	testuc "github.com/mrpandey/gobp/src/internal/core/usecase/testutil"
	"github.com/mrpandey/gobp/src/util"
)

func SetupRESTServer(t *testing.T) (*resth.RestServer, *usecase.UseCases) {
	cfg := util.NewConfig()
	disableCallerLog := false
	logger := util.NewLogger(cfg.LogLevel, cfg.DGN, disableCallerLog)

	useCases := testuc.NewTestUseCase(t)
	validator := util.NewRequestBodyValidator(logger)

	return resth.NewRestServer(logger, validator, useCases), useCases
}

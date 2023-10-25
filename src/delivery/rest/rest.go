package restd

import (
	"context"
	"fmt"
	"net/http"
	"time"

	resth "github.com/mrpandey/gobp/src/delivery/rest/handlers"
	usecase "github.com/mrpandey/gobp/src/internal/core/usecase"
	"github.com/mrpandey/gobp/src/util"
)

func InitHTTPDelivery(
	_ context.Context,
	cfg *util.Config,
	logger *util.StandardLogger,
	validator *util.Validator,
	useCases *usecase.UseCases,
) {
	address := fmt.Sprintf("%s:%d", cfg.Host, cfg.RESTPort)
	gobpHandler := resth.NewRestServer(logger, validator, useCases)
	router := NewRouter(cfg, logger, gobpHandler, useCases.Auth)

	httpServer := &http.Server{
		Addr:              address,
		ReadHeaderTimeout: 1 * time.Second,
		Handler:           router,
	}

	logger.Infof("HTTP server running on %s", httpServer.Addr)
	err := httpServer.ListenAndServe()
	if err != nil {
		logger.Panic(err)
	}
}

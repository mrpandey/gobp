package restd

import (
	resth "github.com/mrpandey/gobp/src/delivery/rest/handlers"
	authdom "github.com/mrpandey/gobp/src/internal/core/domain/auth"
	"github.com/mrpandey/gobp/src/util"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewRouter(
	_ *util.Config,
	_ *util.StandardLogger,
	s *resth.RestServer,
	_ authdom.AuthUseCaseInterface,
) *chi.Mux {
	router := chi.NewRouter()

	basicRouter := chi.NewRouter()
	basicRouter.Get("/ping", s.PingHandler)
	basicRouter.Get("/health", s.HealthHandler)
	router.Mount("/gobp", basicRouter)

	v1Router := chi.NewRouter()

	// middlewares
	v1Router.Use(middleware.AllowContentType("application/json"))
	// v1Router.Use(VerifyAccessToken(logger, auth, cfg.SecretKey))
	v1Router.Use(middleware.Logger)

	v1Router.Get("/furniture/{id}", s.GetFurniture)
	v1Router.Post("/furniture", s.AddFurniture)
	// basicRouter.Post("/auth/token", s.CreateToken)

	router.Mount("/gobp/v1", v1Router)

	return router
}

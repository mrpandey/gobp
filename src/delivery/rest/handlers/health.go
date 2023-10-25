package resth

import (
	"net/http"

	healthdom "github.com/mrpandey/gobp/src/internal/core/domain/health"
)

func (s *RestServer) PingHandler(w http.ResponseWriter, _ *http.Request) {
	pingMsg := s.useCases.Health.Ping()
	response := healthdom.Ping{
		Message: pingMsg,
	}
	sendJSONResponse(s.logger, w, http.StatusOK, response)
}

func (s *RestServer) HealthHandler(w http.ResponseWriter, r *http.Request) {
	dbHealth := s.useCases.Health.HealthCheck(r.Context())

	responseCode := http.StatusOK
	if !dbHealth.PostgresReachable {
		responseCode = http.StatusServiceUnavailable
	}
	sendJSONResponse(s.logger, w, responseCode, dbHealth)
}

package resth

import (
	"fmt"
	"net/http"

	authdom "github.com/mrpandey/gobp/src/internal/core/domain/auth"
	gobperror "github.com/mrpandey/gobp/src/internal/core/domain/error"
)

func (s *RestServer) CreateToken(w http.ResponseWriter, r *http.Request) {
	clientSlug, clientSecret, ok := r.BasicAuth()
	if !ok {
		w.Header().Add("WWW-Authenticate", fmt.Sprintf(`Basic realm="%v", charset="UTF-8"`, authdom.ClientCredsRealm))
		SendErrorResponse(s.logger, w, gobperror.ErrAuthHeaderMissing)
		return
	}

	tokenResponse, err := s.useCases.Auth.CreateToken(r.Context(), clientSlug, clientSecret)
	if err != nil {
		w.Header().Add("WWW-Authenticate", fmt.Sprintf(`Basic realm="%v", charset="UTF-8"`, authdom.ClientCredsRealm))
		SendErrorResponse(s.logger, w, err)
		return
	}

	sendJSONResponse(s.logger, w, http.StatusOK, tokenResponse)
}

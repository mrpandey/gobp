package restd

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	resth "github.com/mrpandey/gobp/src/delivery/rest/handlers"
	authdom "github.com/mrpandey/gobp/src/internal/core/domain/auth"
	cdom "github.com/mrpandey/gobp/src/internal/core/domain/common"
	gobperror "github.com/mrpandey/gobp/src/internal/core/domain/error"
	"github.com/mrpandey/gobp/src/util"
)

func VerifyAccessToken(
	logger *util.StandardLogger,
	auth authdom.AuthUseCaseInterface,
	signingKey string,
) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			reqToken := r.Header.Get("Authorization")
			if !strings.HasPrefix(reqToken, "Bearer ") {
				sendAuthErrorResponse(logger, w, gobperror.ErrAuthHeaderMissing)
				return
			}

			accessToken := reqToken[len("Bearer "):]

			claims, ok := auth.VerifyToken(accessToken, signingKey, authdom.Access)
			if !ok {
				sendAuthErrorResponse(logger, w, gobperror.ErrInvalidToken)
				return
			}

			// set client id in context for use by handlers
			ctx := context.WithValue(r.Context(), cdom.Source, claims.Subject)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}

func sendAuthErrorResponse(logger *util.StandardLogger, w http.ResponseWriter, err error) {
	w.Header().Add("WWW-Authenticate", fmt.Sprintf(`Bearer realm="%v", charset="UTF-8"`, authdom.ClientBearerRealm))
	resth.SendErrorResponse(logger, w, err)
}

package authdom

import (
	"context"
	"time"

	cdom "github.com/mrpandey/gobp/src/internal/core/domain/common"

	"github.com/golang-jwt/jwt/v4"
)

type AuthRealm string

type JWTType string

const (
	ClientCredsRealm  AuthRealm = "client-credentials"
	ClientBearerRealm AuthRealm = "bearer-v1"

	Access JWTType = "access"

	TokenIssuer     = "gobp"
	AccessTokenLife = 15 * time.Minute
)

type JWTClaims struct {
	jwt.RegisteredClaims
	Type JWTType `json:"type"` // access or refresh
}

type AuthUseCaseInterface interface {
	CreateToken(ctx context.Context, clientSlug, clientSecret string) (TokenResponse, error)
	VerifyToken(accessToken, signingKey string, tokenType JWTType) (JWTClaims, bool)
}

type AuthRepoInterface interface {
	WithTx(tm cdom.TxnManagerInterface) AuthRepoInterface
	GetCreds(ctx context.Context, clientSlug string) (CredRecord, error)
}

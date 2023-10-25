package authuc

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v4"
	authdom "github.com/mrpandey/gobp/src/internal/core/domain/auth"
	"github.com/mrpandey/gobp/src/util"
)

const (
	BcryptWorkFactor int = 8
)

func verifyClaims(claims authdom.JWTClaims, tokenType authdom.JWTType) bool {
	now := time.Now()

	if !claims.VerifyExpiresAt(now, true) {
		return false
	}

	if !claims.VerifyIssuer(authdom.TokenIssuer, true) {
		return false
	}

	if claims.Type != tokenType {
		return false
	}

	return true
}

func newTokenResponse(
	ctx context.Context,
	logger *util.StandardLogger,
	clientSlug, accessSigningKey string,
) (authdom.TokenResponse, error) {
	var emptyTokenResponse authdom.TokenResponse

	accessExpiry := time.Now().Add(authdom.AccessTokenLife)
	accessToken, err := newToken(ctx, logger, clientSlug, authdom.Access, accessSigningKey, accessExpiry)
	if err != nil {
		return emptyTokenResponse, err
	}

	return authdom.TokenResponse{
		AccessToken:     accessToken,
		AccessExpiresIn: uint(authdom.AccessTokenLife / time.Second),
		TokenType:       "Bearer",
	}, nil
}

// Generates a new JWT token for the given subject/client.
func newToken(
	_ context.Context,
	logger *util.StandardLogger,
	subject string,
	tokenType authdom.JWTType,
	signingKey string,
	expiry time.Time,
) (string, error) {
	claims := authdom.JWTClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiry),
			Issuer:    authdom.TokenIssuer,
			Subject:   subject,
		},
		Type: tokenType,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenString, err := token.SignedString([]byte(signingKey))

	if err != nil {
		logger.Errorf("Could not generate JWT token. %v", err)

		return "", err
	}

	return tokenString, nil
}

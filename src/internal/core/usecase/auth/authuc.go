package authuc

import (
	"context"
	"encoding/base64"
	"errors"
	"strings"

	authdom "github.com/mrpandey/gobp/src/internal/core/domain/auth"
	cdom "github.com/mrpandey/gobp/src/internal/core/domain/common"
	gobperror "github.com/mrpandey/gobp/src/internal/core/domain/error"
	"github.com/mrpandey/gobp/src/util"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// Implements AuthUseCaseInterface.
type authUseCase struct {
	cfg      *util.Config
	logger   *util.StandardLogger
	txnMgr   cdom.TxnManagerInterface
	authRepo authdom.AuthRepoInterface
}

func NewAuthUseCase(
	cfg *util.Config,
	logger *util.StandardLogger,
	txnMgr cdom.TxnManagerInterface,
	authRepo authdom.AuthRepoInterface,
) *authUseCase {
	return &authUseCase{
		cfg:      cfg,
		logger:   logger,
		txnMgr:   txnMgr,
		authRepo: authRepo,
	}
}

func (uc *authUseCase) CreateToken(
	ctx context.Context,
	clientSlug string,
	clientSecret string,
) (authdom.TokenResponse, error) {
	emptyTokenResponse := authdom.TokenResponse{}

	// verify supplied credentials
	creds, err := uc.authRepo.GetCreds(ctx, clientSlug)
	if err != nil {
		if errors.Is(err, gobperror.ErrRecordNotFound) {
			return emptyTokenResponse, gobperror.ErrWrongCreds
		}
		return emptyTokenResponse, err
	}

	suppliedSecret, err := base64.RawURLEncoding.DecodeString(clientSecret)
	if err != nil {
		return emptyTokenResponse, gobperror.ErrWrongCreds
	}
	hashedSecretInDB, err := base64.RawURLEncoding.DecodeString(creds.HashedSecret)
	if err != nil {

		uc.logger.Errorf("Could not decode hashed secret in database for clientSlug=%v. %v", clientSlug, err)
		return emptyTokenResponse, gobperror.ErrUnexpected
	}

	err = bcrypt.CompareHashAndPassword(hashedSecretInDB, suppliedSecret)
	if err != nil {
		return emptyTokenResponse, gobperror.ErrWrongCreds
	}

	// block-check should be after verifying the secret
	if creds.IsBlocked {
		return emptyTokenResponse, gobperror.ErrBlocked
	}

	accessSigningKey := uc.cfg.SecretKey
	tokenResponse, err := newTokenResponse(
		ctx,
		uc.logger,
		clientSlug,
		accessSigningKey,
	)
	if err != nil {
		return emptyTokenResponse, nil
	}

	return tokenResponse, err
}

// Returns a false boolean if the token is not valid. Otherwise returns a true boolean along with JWT claims.
func (au *authUseCase) VerifyToken(
	accessToken string,
	signingKey string,
	tokenType authdom.JWTType,
) (authdom.JWTClaims, bool) {
	var claims authdom.JWTClaims

	jwtParser := jwt.NewParser()
	token, tokenParts, err := jwtParser.ParseUnverified(accessToken, &claims)
	if err != nil {
		return authdom.JWTClaims{}, false
	}

	if !verifyClaims(claims, tokenType) {
		return authdom.JWTClaims{}, false
	}

	// verify signing algo
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return authdom.JWTClaims{}, false
	}

	// validate token signature
	err = token.Method.Verify(strings.Join(tokenParts[0:2], "."), tokenParts[2], []byte(signingKey))
	if err != nil {
		return authdom.JWTClaims{}, false
	}

	return claims, true
}

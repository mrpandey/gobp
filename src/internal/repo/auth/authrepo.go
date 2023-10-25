package authrepo

import (
	"context"
	"errors"

	authdom "github.com/mrpandey/gobp/src/internal/core/domain/auth"
	cdom "github.com/mrpandey/gobp/src/internal/core/domain/common"
	gobperror "github.com/mrpandey/gobp/src/internal/core/domain/error"
	"github.com/mrpandey/gobp/src/util"

	"gorm.io/gorm"
)

type authRepo struct {
	db     *gorm.DB
	logger *util.StandardLogger
}

func NewAuthRepo(logger *util.StandardLogger, db *gorm.DB) *authRepo {
	return &authRepo{
		db:     db,
		logger: logger,
	}
}

func (ar *authRepo) WithTx(tm cdom.TxnManagerInterface) authdom.AuthRepoInterface {
	return &authRepo{
		db:     tm.GetTx(),
		logger: ar.logger,
	}
}

func (ar *authRepo) GetCreds(ctx context.Context, clientSlug string) (authdom.CredRecord, error) {
	var clientCreds ClientCred
	var emptyRecord authdom.CredRecord

	err := ar.db.WithContext(ctx).
		Select("id", "slug", "hashed_secret", "is_blocked").
		Where("slug = ?", clientSlug).
		First(&clientCreds).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return emptyRecord, gobperror.ErrRecordNotFound
		}

		ar.logger.Errorf("Failed to get creds from database for client=%v. %v", clientSlug, err)
		return emptyRecord, err
	}

	return authdom.CredRecord{
		ID:           clientCreds.ID,
		Slug:         clientCreds.Slug,
		HashedSecret: clientCreds.HashedSecret,
		IsBlocked:    clientCreds.IsBlocked,
	}, nil
}

package db

import (
	"oauth_api/src/domain/access_token"
	restError "oauth_api/src/utils/errors"
)

func NewRepository() DbRepo {
	return &dbRepository{}
}

type DbRepo interface {
	GetById(string) (*access_token.AccessToken, *restError.RestError)
}

type dbRepository struct {

}

func (r *dbRepository) GetById(accessTokenId string) (*access_token.AccessToken, *restError.RestError) {
	return nil, restError.NewInternalServerError("database connection not implemented yet")
}
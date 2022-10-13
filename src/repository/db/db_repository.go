package db

import (
	"oauth_api/src/client/cassandra"
	"oauth_api/src/domain/access_token"
	restError "oauth_api/src/utils/errors"
)

const (
	queryGetAccessToken="select access_token, user_id, client_id, expires from access_tokens where access_token=?;"
	queryInsertAccessToken="insert into access_tokens(access_token, user_id, client_id, expires) values(?,?,?,?); "
	queryUpdateExpirationAccessToken="update access_tokens set expires=? where access_token=?;"

	errNotFound="not found"
)

func NewRepository() DbRepo {
	return &dbRepository{}
}

type DbRepo interface {
	GetById(string) (*access_token.AccessToken, *restError.RestError)
	Create(access_token.AccessToken) *restError.RestError
	UpdateExpirationTime(access_token.AccessToken) *restError.RestError
}

type dbRepository struct {

}

func (r *dbRepository) GetById(accessTokenId string) (*access_token.AccessToken, *restError.RestError) {
	session, err := cassandra.GetSession()
	if err != nil {
		return nil, restError.NewInternalServerError(err.Error())
	}
	defer session.Close()

	var result access_token.AccessToken
	if err := session.Query(queryGetAccessToken, accessTokenId).Scan(
		&result.AccessToken, &result.UserId, &result.ClientId, &result.Expires,
	); err != nil {
		if err.Error() == errNotFound {
			return nil, restError.NewNotFoundError("no access token found with given id")
		}

		return nil, restError.NewInternalServerError(err.Error())
	}

	return &result,nil
}

func (r *dbRepository) Create(at access_token.AccessToken) *restError.RestError {
	session, err := cassandra.GetSession()
	if err != nil {
		return restError.NewInternalServerError(err.Error())
	}
	defer session.Close()

	if err := session.Query(
		queryInsertAccessToken, at.AccessToken, at.UserId, at.ClientId, at.Expires,
	).Exec(); err != nil {
		return restError.NewInternalServerError(err.Error())
	}

	return nil
}

func (r *dbRepository) UpdateExpirationTime(at access_token.AccessToken) *restError.RestError {
	session, err := cassandra.GetSession()
	if err != nil {
		return restError.NewInternalServerError(err.Error())
	}
	defer session.Close()

	if err := session.Query(
		queryUpdateExpirationAccessToken,at.Expires, at.AccessToken,
	).Exec(); err != nil {
		return restError.NewInternalServerError(err.Error())
	}

	return nil
}
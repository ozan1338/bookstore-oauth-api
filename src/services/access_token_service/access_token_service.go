package access_token_service

import (
	"oauth_api/src/domain/access_token"
	"oauth_api/src/repository/db"
	"oauth_api/src/repository/rest"
	restError "oauth_api/src/utils/errors"
	"strings"
)

type Service interface {
	GetById(string) (*access_token.AccessToken, restError.RestError)
	Create(access_token.AccessTokenRequest) (*access_token.AccessToken,restError.RestError)
	UpdateExpirationTime(access_token.AccessToken) restError.RestError
}

type service struct {
	restUserRepo rest.RestUsersRepository
	dbRepo db.DbRepo
}

func NewService(restUserRepo rest.RestUsersRepository,dbRepo db.DbRepo) Service {
	return &service{
		restUserRepo: restUserRepo,
		dbRepo: dbRepo,
	}
}

func (s *service) GetById(accessTokenId string) (*access_token.AccessToken, restError.RestError) {
	accessTokenId = strings.TrimSpace(accessTokenId)

	if len(accessTokenId) == 0 {
		return nil, restError.NewBadRequestError("invalid access token id")
	}

	accessToken, err := s.dbRepo.GetById(accessTokenId)

	if err != nil {
		return nil,err
	}
	return accessToken, nil
}

func (s *service) Create(request access_token.AccessTokenRequest) (*access_token.AccessToken,restError.RestError) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	//TODO: Support both client credential and password grant type

	//Authenticate the user againt the user api
	user, err := s.restUserRepo.Login(request.Username, request.Password)
	if err != nil {
		return nil,err
	}

	//Generate a new access token
	at := access_token.GetNewAccessToken(user.Id)
	at.Generate()

	//Save the new access token in Cassandra
	if err := s.dbRepo.Create(at); err != nil {
		return nil,err
	}

	return &at,nil
}

func (s *service) UpdateExpirationTime(at access_token.AccessToken) restError.RestError {
	if err := at.Validate(); err != nil {
		return err
	}


	return s.dbRepo.UpdateExpirationTime(at)
}
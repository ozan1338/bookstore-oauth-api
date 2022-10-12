package access_token

import (
	restError "oauth_api/src/utils/errors"
	"strings"
)

type Repository interface{
	GetById(string) (*AccessToken, *restError.RestError)
}

type Service interface {
	GetById(string) (*AccessToken, *restError.RestError)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(accessTokenId string) (*AccessToken, *restError.RestError) {
	accessTokenId = strings.TrimSpace(accessTokenId)

	if len(accessTokenId) == 0 {
		return nil, restError.NewBadRequestError("invalid access token id")
	}

	accessToken, err := s.repository.GetById(accessTokenId)

	if err != nil {
		return nil,err
	}
	return accessToken, nil
}
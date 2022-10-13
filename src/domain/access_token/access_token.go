package access_token

import (
	"strings"
	"time"

	resError "oauth_api/src/utils/errors"
)

const (
	expirationTime = 24
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func (at *AccessToken) Validate() *resError.RestError {
	at.AccessToken = strings.TrimSpace(at.AccessToken)

	if at.AccessToken == "" {
		return resError.NewBadRequestError("invalid access token id")
	}

	if at.UserId <= 0 {
		return resError.NewBadRequestError("invalid user id")
	}

	if at.ClientId <= 0 {
		return resError.NewBadRequestError("invalid client id")
	}

	if at.Expires <= 0 {
		return resError.NewBadRequestError("invalid expiration id")
	}

	return nil
}

func GetNewAccessToken() *AccessToken {
	return &AccessToken{
		Expires: time.Now().UTC().Add(expirationTime*time.Hour).Unix(),
	}
}

func(at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now())
}
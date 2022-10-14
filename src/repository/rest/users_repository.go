package rest

import (
	"encoding/json"
	"oauth_api/src/domain/users"
	restError "oauth_api/src/utils/errors"
	"time"

	resty "github.com/go-resty/resty/v2"
	"github.com/mercadolibre/golang-restclient/rest"
)

var (
	userRestClient = rest.RequestBuilder{
		BaseURL: "http://localhost:8080",
		Timeout: 100 * time.Millisecond,
	}

	// client = resty.Client{
	// 	BaseURL: "http://localhost",
	// }
	client = resty.Client{
		BaseURL: "http://localhost:8080",
	}
)


type RestUsersRepository interface{
	Login(string,string) (*users.User, *restError.RestError)
}

type usersRepository struct {}

func NewRestUsersRepository() RestUsersRepository {
	return &usersRepository{}
}

func (r *usersRepository) Login(email string, password string) (*users.User, *restError.RestError) {
	request := users.UserLoginRequest{
		Email: email,
		Password: password,
	}


	response := userRestClient.Post("/users/login", request)
	
	if response == nil || response.Response == nil {
		return nil, restError.NewInternalServerError("invalid rest client response when to try login user")
	}

	if response.StatusCode > 399 {
		// fmt.Println("HERERERE")
		var respError restError.RestError
		if err := json.Unmarshal(response.Bytes(), &respError); err != nil {
			return nil, restError.NewInternalServerError("invalid err interface when trying to logging user")
		}
		return nil ,&respError
	}

	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, restError.NewInternalServerError("error when trying to unmarshall user response")
	}
	return &user,nil
}
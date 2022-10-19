package rest

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
)

// var (
// 	testClient = resty.New()
// )

func TestMain(m *testing.M) {
	log.Println("STARTING TEST")
	// httpmock.ActivateNonDefault(client.GetClient())
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestLoginUserTimeoutFromApi(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "http://localhost:8080/users/login",
		ReqBody:      `{"email":"email@gmail.com","password":"the-password"}`,
		RespHTTPCode: -1,
		RespBody:     `{}`,
	})

	repository := usersRepository{}

	user, err := repository.Login("email@gmail.com", "the-password")
	
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.GetStatus())
	assert.EqualValues(t, "invalid rest client response when to try login user", err.GetMessage())
}


func TestLoginUserInvalidErrorInterface(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "http://localhost:8080/users/login",
		ReqBody:      `{"email":"email@gmail.com","password":"the-password"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": "invalid login credentials", "status": "404", "error": "not_found"}`,
	})

	repository := usersRepository{}

	user, err := repository.Login("email@gmail.com","the-password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.GetStatus())
	assert.EqualValues(t, "invalid err interface when trying to logging user", err.GetMessage())
}

func TestLoginUserInvalidLoginCredentials(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod: http.MethodPost,
		URL: "http://localhost:8080/users/login",
		ReqBody: `{"email":"email@gmail.com","password":"the-password"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody: `{"message": "invalid login credentials", "status": 404, "error": "not_found"}`,
	})

	repository := usersRepository{}

	user, err := repository.Login("email@gmail.com","the-password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.GetStatus())
	assert.EqualValues(t, "invalid login credentials", err.GetMessage())
}

func TestLoginUserInvalidUserJsonResponse(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod: http.MethodPost,
		URL: "http://localhost:8080/users/login",
		ReqBody: `{"email":"email@gmail.com","password":"the-password"}`,
		RespHTTPCode: http.StatusOK,
		RespBody: `{"id":12,"first_name":Ozan,"last_name":"Hehe","email":"me@here.com","date_created":"2022-10-11 03:52:04","status":"active"}`,
	})

	repository := usersRepository{}

	user, err := repository.Login("email@gmail.com","the-password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.GetStatus())
	assert.EqualValues(t, "error when trying to unmarshall user response", err.GetMessage())
}

func TestLoginUserNoError(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod: http.MethodPost,
		URL: "http://localhost:8080/users/login",
		ReqBody: `{"email":"email@gmail.com","password":"the-password"}`,
		RespHTTPCode: http.StatusOK,
		RespBody: `{"id":12,"first_name":"Ozan","last_name":"Hehe","email":"me@here.com","date_created":"2022-10-11 03:52:04","status":"active"}`,
	})

	repository := usersRepository{}

	user, err := repository.Login("email@gmail.com","the-password")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 12, user.Id)
	assert.EqualValues(t, "Ozan", user.FirstName)
	assert.EqualValues(t, "Hehe", user.LastName)
	assert.EqualValues(t, "me@here.com", user.Email)
}
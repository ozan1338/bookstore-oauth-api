package http

import (
	"net/http"
	atDomain "oauth_api/src/domain/access_token"
	"oauth_api/src/services/access_token_service"
	restError "oauth_api/src/utils/errors"

	"github.com/gin-gonic/gin"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
	Create(*gin.Context)

}

type accessTokenHandler struct {
	service access_token_service.Service
}

func NewHandler(service access_token_service.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (handler *accessTokenHandler) GetById(c *gin.Context) {
	accessToken, err := handler.service.GetById(c.Param("access_token_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, accessToken)
}

func (handler *accessTokenHandler) Create(c *gin.Context) {
	var request atDomain.AccessTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		// fmt.Println(err)
		restError := restError.NewBadRequestError("invalid json body")
		c.JSON(restError.Status, restError)
		return
	}

	at,err := handler.service.Create(request); 
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusCreated, at)
}
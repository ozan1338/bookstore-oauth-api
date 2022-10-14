package app

import (
	"oauth_api/src/http"
	"oauth_api/src/repository/db"
	"oauth_api/src/repository/rest"
	"oauth_api/src/services/access_token_service"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	// session, dbErr := cassandra.GetSession()
	// if dbErr != nil {
	// 	panic(dbErr)
	// }
	// defer session.Close()

	atService := access_token_service.NewService(rest.NewRestUsersRepository(),db.NewRepository())
	atHandler := http.NewHandler(atService)

	router.POST("/oauth/access_token", atHandler.Create)
	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.Run(":8081")
}
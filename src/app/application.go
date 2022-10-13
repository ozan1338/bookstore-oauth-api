package app

import (
	"oauth_api/src/client/cassandra"
	"oauth_api/src/domain/access_token"
	"oauth_api/src/http"
	"oauth_api/src/repository/db"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	session, dbErr := cassandra.GetSession()
	if dbErr != nil {
		panic(dbErr)
	}
	defer session.Close()

	atService := access_token.NewService(db.NewRepository())
	atHandler := http.NewHandler(atService)

	router.POST("/oauth/access_token", atHandler.Create)
	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.Run(":8081")
}
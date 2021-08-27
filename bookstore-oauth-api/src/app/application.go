package app

import (
	"github.com/gin-gonic/gin"

	"bookstore-oauth-api/src/repository/db"
	"bookstore-oauth-api/src/repository/rest"
	atService "bookstore-oauth-api/src/services/access_token"
	"bookstore-oauth-api/src/http"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atHandler := http.NewAccessTokenHandler(atService.NewService(rest.NewRestUsersRepository(), db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)

	router.Run(":8080")
}
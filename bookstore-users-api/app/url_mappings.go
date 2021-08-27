package app

import (
	"bookstore-users-api/controllers/ping"
	"bookstore-users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping_controller.Ping)

	router.POST("/users", users_controller.Create)
	router.GET("/users/:user_id", users_controller.Get)
	router.PUT("/users/:user_id", users_controller.Update)
	router.PATCH("/users/:user_id", users_controller.Update)
	router.DELETE("/users/:user_id", users_controller.Delete)
	router.GET("/internal/users/search", users_controller.Search)
	router.POST("/users/login", users_controller.Login)
}
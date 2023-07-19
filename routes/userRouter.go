package routes

import (
	"github.com/gin-gonic/gin"
	"jwt-auth/controller"
	"jwt-auth/middleware"
)

func UserRoutes(router *gin.Engine) {
	router.Use(middleware.Auth)
	router.GET("/users", controller.GetUser)
	router.GET("/users/:user_id", controller.GetUserById)
}

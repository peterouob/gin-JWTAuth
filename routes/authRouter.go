package routes

import (
	"github.com/gin-gonic/gin"
	"jwt-auth/controller"
)

func AuthRoutes(router *gin.Engine) {
	router.POST("/users/signup", controller.Signup)
	router.POST("/users/login", controller.Login)
}

package main

import (
	"github.com/gin-gonic/gin"
	"jwt-auth/config"
	"jwt-auth/routes"
	"net/http"
)

var port string

func main() {
	port = config.Config.GetString("server.port")
	if port == "" {
		port = "8084"
	}
	r := gin.New()
	r.Use(gin.Logger())

	routes.AuthRoutes(r)
	routes.UserRoutes(r)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "Hello",
		})
	})
	r.Run(":" + port)
}

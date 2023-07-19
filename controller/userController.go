package controller

import (
	"github.com/gin-gonic/gin"
	validator2 "github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
	"jwt-auth/database"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "users")
var validator = validator2.New()

func HashPassword() {}

func VerifyPassword() {}

func Login(c *gin.Context) {}

func Signup(c *gin.Context) {}

func GetUser(c *gin.Context) {}

func GetUserById(c *gin.Context) {
	userID := c.Param("user_id")
}

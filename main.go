package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/sithsithsith/cybe-auth/core/lib/logger"
	user_service "github.com/sithsithsith/cybe-auth/http/services"
)

func TestLogger() gin.HandlerFunc {
	errLogger := logger.NewLogger(logrus.ErrorLevel)

	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"message": "Internal Server Error",
				})
				errLogger.Fatal(err)
			}
		}()
		ctx.Next()
	}
}

func GetUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, user_service.GetUsersList())
}

func CreateUser(c *gin.Context) {
	var newUser user_service.User
	//built-in bindJson
	if err := c.BindJSON(user_service.NewUser()); err != nil {
		c.JSON(http.StatusBadRequest, "Failed to")
		// c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	u, err := user_service.CreateUser(newUser)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.IndentedJSON(http.StatusCreated, u)
}

func main() {
	router := gin.Default()
	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&logrus.JSONFormatter{})
	router.Use(TestLogger())
	router.GET("users", GetUsers)
	router.POST("user", CreateUser)
	router.Run("localhost:5200")
}

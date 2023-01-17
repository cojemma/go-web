package main

import (
	"go-web/user"
	"go-web/utils"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	env := os.Getenv("ENVIROMENT")
	if env == "dev" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.Use(
		utils.CallApiCount,
	)

	router.GET("", user.GetUsers)
	router.GET("/users", user.GetUsers)
	router.POST("/users", user.PostUser)
	router.GET("/users/:id", user.GetUserByID)
	router.DELETE("users/:id", user.DeleteUser)
	router.PUT("/users/:id", user.UpdateUser)

	router.Run()
}

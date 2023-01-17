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

	userApi := router.Group("/users")
	userApi.GET("", user.GetUsers)
	userApi.POST("", user.PostUser)
	userApi.GET("/:id", user.GetUserByID)
	userApi.DELETE("/:id", user.DeleteUser)
	userApi.PUT("/:id", user.UpdateUser)

	router.Run()
}

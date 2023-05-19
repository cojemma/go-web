package main

import (
	"go-web/api"
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

	router.GET("", api.GetUsers)

	userApi := router.Group("/users")
	userApi.GET("", api.GetUsers)
	userApi.POST("", api.PostUser)
	userApi.GET("/:id", api.GetUserByID)
	userApi.DELETE("/:id", api.DeleteUser)
	userApi.PUT("/:id", api.UpdateUser)

	router.Run(":8080")
}

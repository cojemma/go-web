package api

import (
	"go-web/database"
	"go-web/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var sqldb *gorm.DB

func init() {
	sqldb = database.GetDB()
	sqldb.AutoMigrate(&model.User{})
}

func GetUsers(c *gin.Context) {
	users, err := model.GetUsers()
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, users)
}

func PostUser(c *gin.Context) {
	var newUser model.User
	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	err := model.CreateUser(&newUser)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusCreated, newUser)
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	userID, _ := strconv.Atoi(id)

	user, err := model.GetUser(&model.UserScope{
		ID: userID,
	})

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	userID, _ := strconv.Atoi(id)

	err := model.DeleteUser(&model.UserScope{
		ID: userID,
	})

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Delete success"})
	}
}

func UpdateUser(c *gin.Context) {
	var newuser model.User
	c.ShouldBind(&newuser)

	err := model.UpdateUser(&newuser)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
	} else {
		c.JSON(http.StatusOK, newuser)
	}
}

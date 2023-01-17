package user

import (
	"go-web/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var sqldb *gorm.DB

func init() {
	sqldb = database.GetDB()
	sqldb.AutoMigrate(&User{})
}

func GetUsers(c *gin.Context) {
	var users []User
	sqldb.Find(&users)
	c.IndentedJSON(http.StatusOK, users)
}

func PostUser(c *gin.Context) {
	var newUser User
	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	sqldb.Create(&newUser)

	c.IndentedJSON(http.StatusCreated, newUser)
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")

	var users []User
	result := sqldb.Where("id = ?", id).Find(&users)

	if result.RowsAffected == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
	} else {
		c.IndentedJSON(http.StatusOK, users)
	}
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	result := sqldb.Delete(&User{}, id)

	if result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Delete success"})
	}
}

func UpdateUser(c *gin.Context) {
	var newuser User
	c.ShouldBind(&newuser)

	var user User
	result := sqldb.First(&user, c.Param("id"))
	if result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
	} else {
		user.UserName = newuser.UserName
		sqldb.Save(&user)
		c.IndentedJSON(http.StatusOK, user)
	}
}

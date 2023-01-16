package user

import (
	"fmt"
	"go-web/database"
	"go-web/utils"
	"net/http"
	"os/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var sqldb *gorm.DB

func init() {
	sqldb = database.GetDB()
	sqldb.AutoMigrate(&User{})
}

func GetUsers(c *gin.Context) {
	count := utils.CallApiCount("main")
	fmt.Printf("getall Has been call for %v times\n", count)

	var users []user.User
	sqldb.Find(&users)
	c.IndentedJSON(http.StatusOK, users)
}

func PostUser(c *gin.Context) {
	utils.CallApiCount("postnew")

	var newUser User
	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	sqldb.Create(&newUser)

	c.IndentedJSON(http.StatusCreated, newUser)
}

func GetUserByID(c *gin.Context) {
	utils.CallApiCount("getby")

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
	utils.CallApiCount("delete")

	id := c.Param("id")

	result := sqldb.Delete(&User{}, id)

	if result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Delete success"})
	}
}

func UpdateUser(c *gin.Context) {
	utils.CallApiCount("update")

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

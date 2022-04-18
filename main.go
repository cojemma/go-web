package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       string
	UserName string
}

var sqldb, _ = gorm.Open(mysql.Open("root:db@tcp(db:3306)/mysql?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})

var rdb = redis.NewClient(&redis.Options{
	Addr:     "redis:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func getUsers(c *gin.Context) {
	count := callApiCount("main")
	fmt.Printf("getall Has been call for %v times\n", count)

	var users []User
	sqldb.Find(&users)
	c.IndentedJSON(http.StatusOK, users)
}

func postUser(c *gin.Context) {
	callApiCount("postnew")

	var newUser User
	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	sqldb.Create(&newUser)

	c.IndentedJSON(http.StatusCreated, newUser)
}

func getUserByID(c *gin.Context) {
	callApiCount("getby")

	id := c.Param("id")

	var users []User
	result := sqldb.Where("id = ?", id).Find(&users)

	if result.RowsAffected == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
	} else {
		c.IndentedJSON(http.StatusOK, users)
	}
}

func deleteUser(c *gin.Context) {
	callApiCount("delete")

	id := c.Param("id")

	result := sqldb.Delete(&User{}, id)

	if result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Delete success"})
	}
}

func updateUser(c *gin.Context) {
	callApiCount("update")

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

func callApiCount(api string) int {
	count := 1
	if val, err := rdb.Get(api).Result(); err != redis.Nil {
		if val, errt := strconv.Atoi(val); errt == nil {
			count += val
		}
	} else {
		fmt.Println("first call")
	}
	rdb.Set(api, count, 0)
	return count
}

func main() {
	_, err := gorm.Open(mysql.Open("root:db@tcp(db:3306)/mysql?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	sqldb.AutoMigrate(&User{})

	router := gin.Default()
	router.GET("", getUsers)
	router.GET("/users", getUsers)
	router.POST("/users", postUser)
	router.GET("/users/:id", getUserByID)
	router.DELETE("users/:id", deleteUser)
	router.PUT("/users/:id", updateUser)

	router.Run()

}

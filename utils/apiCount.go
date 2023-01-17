package utils

import (
	"fmt"
	"go-web/database"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

var rdb *redis.Client

func init() {
	rdb = database.GetCache()
}

func CallApiCount(c *gin.Context) {
	api := c.Request.URL.Path

	count := 1
	if val, err := rdb.Get(api).Result(); err != redis.Nil {
		if val, errt := strconv.Atoi(val); errt == nil {
			count += val
		}
		fmt.Printf("API %s count: %v\n", api, count)
	} else {
		fmt.Println("first call")
	}
	rdb.Set(api, count, 0)

	c.Next()
}

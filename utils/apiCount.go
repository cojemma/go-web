package utils

import (
	"fmt"
	"go-web/database"
	"strconv"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func init() {
	rdb = database.GetCache()
}

func CallApiCount(api string) int {
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

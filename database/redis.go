package database

import "github.com/go-redis/redis"

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func connectRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if rdb.Ping().Err() != nil {
		rdb.Close()
		rdb = nil
	}
}

func GetCache() *redis.Client {
	if rdb == nil {
		connectRedis()
	}

	return rdb
}

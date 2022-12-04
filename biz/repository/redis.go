package repository

import (
	"github.com/go-redis/redis"
	"github.com/lutasam/check_in_sys/biz/utils"
)

var redisDB *redis.Client

func init() {
	redisDB = redis.NewClient(&redis.Options{
		Addr:     utils.GetConfigString("redis.address"),
		Password: utils.GetConfigString("redis.password"),
	})
	_, err := redisDB.Ping().Result()
	if err != nil {
		panic(err)
	}
}

func GetRedis() *redis.Client {
	return redisDB
}

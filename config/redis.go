package config

import (
	"ginDemo/global"
	"github.com/go-redis/redis"
	"log"
)

func InitRedis() {
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		DB:       0,
		Password: "",
	})
	_, err := RedisClient.Ping().Result()
	if err != nil {
		log.Fatal("连接Redis异常,请重试!!!", err)
	}
	global.Redis = RedisClient
}

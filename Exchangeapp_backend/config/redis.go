package config

import (
	"exchangeapp/global"
	"log"

	"github.com/go-redis/redis"
)

func initRedis(){

	addr := AppConfig.Redis.Addr
	db := AppConfig.Redis.DB
	password := AppConfig.Redis.Password 

	RedisClient := redis.NewClient(&redis.Options{
		Addr: addr,
		DB: db,
		Password: password,
	})

	_, err := RedisClient.Ping().Result()

	if err !=nil{
		log.Fatalf("Failed to connect to Redis, got error: %v", err)
	}

	global.RedisDB = RedisClient
}
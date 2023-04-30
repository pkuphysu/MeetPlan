package redis

import (
	"context"
	"fmt"

	"meetplan/config"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func Init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.GetConf().Redis.Address,
		Username: config.GetConf().Redis.Username,
		Password: config.GetConf().Redis.Password,
		DB:       config.GetConf().Redis.DB,
	})
	if err := RedisClient.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
}
func Close() {
	err := RedisClient.Close()
	if err != nil {
		fmt.Printf("close redis error - %v\n", err)
		return
	}
}

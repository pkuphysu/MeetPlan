package redis

import (
	"context"
	"fmt"

	"meetplan/config"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

func Init() {
	Client = redis.NewClient(&redis.Options{
		Addr:     config.GetConf().Redis.Address,
		Username: config.GetConf().Redis.Username,
		Password: config.GetConf().Redis.Password,
		DB:       config.GetConf().Redis.DB,
	})
	if err := Client.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
}
func Close() {
	err := Client.Close()
	if err != nil {
		fmt.Printf("close redis error - %v\n", err)
		return
	}
}

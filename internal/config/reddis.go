package config

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

func NewReddisClient(viper *viper.Viper) *redis.Client {

	rdb := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("REDIS_HOST"),
		Username: viper.GetString("REDIS_USERNAME"),
		Password: viper.GetString("REDIS_PASSWORD"),
		DB:       viper.GetInt("REDIS_DB"),
	})

	pong, err := rdb.Ping(context.Background()).Result()
	fmt.Println(pong, err)

	return rdb
}

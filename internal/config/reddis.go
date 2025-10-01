package config

import "github.com/redis/go-redis/v9"

func NewReddisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // No password set
		DB:       0,  // Use default DB
	})

	return client
}

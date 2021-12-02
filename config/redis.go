package config

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func RedisClient() (*redis.Client, error) {
	redisUrl := GetEnv("REDIS_URL")

	client := redis.NewClient(&redis.Options{
		Addr: redisUrl,
	})

	err := client.Ping(context.Background()).Err()
	return client, err
}

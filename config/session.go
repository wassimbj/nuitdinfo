package config

import (
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis"
)

var store *session.Store

func init() {
	client := redis.New(redis.Config{
		URL: GetEnv("REDIS_URL"),
		// Port: 6379,
		// Host: "app_redis",
	})
	store = session.New(session.Config{
		Storage: client,
	})
}

func Session() *session.Store {
	return store
}

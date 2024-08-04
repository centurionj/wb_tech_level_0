package cache

import (
	"context"
	"github.com/go-redis/redis/v8"

	"wb_tech_level_0/config"
)

// Креатим редис клиент

func NewRedisClient(cfg *config.Config) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: cfg.RedisAddr,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return rdb, nil
}

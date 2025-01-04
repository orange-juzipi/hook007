package cache

import (
	"context"
	"errors"
	"hook007/config"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func Connect() (*redis.Client, error) {
	cfg := config.Get().Redis
	RedisClient = redis.NewClient(&redis.Options{
		Addr:         cfg.Addr,
		Password:     cfg.Pass,
		DB:           cfg.Db,
		MaxRetries:   cfg.MaxRetries,
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
	})

	if err := RedisClient.Ping(context.Background()).Err(); err != nil {
		return nil, errors.New("ping redis err")
	}

	return RedisClient, nil
}

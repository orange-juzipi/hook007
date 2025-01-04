package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type Service interface {
	Login(*gin.Context)
	Register(*gin.Context)
}

type service struct {
	cache *redis.Client
}

func New(cache *redis.Client) Service {
	return &service{
		cache: cache,
	}
}

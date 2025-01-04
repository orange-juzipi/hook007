package middleware

import (
	"hook007/dao/query"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type (
	Middleware interface {
		Token(*gin.Context)
		OpenToken(*gin.Context)
	}

	middleware struct {
		cache *redis.Client
		db    *query.Query
	}
)

func New(redisCache *redis.Client, db *query.Query) Middleware {
	return &middleware{
		cache: redisCache,
		db:    db,
	}
}

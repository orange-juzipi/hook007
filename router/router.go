package router

import (
	"fmt"
	"hook007/config"
	"hook007/dao/query"
	"hook007/pkg/valid"
	"hook007/router/middleware"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"golang.org/x/time/rate"
)

type Server struct {
	db      *query.Query
	cache   *redis.Client
	middles middleware.Middleware
	Mux     *gin.Engine
}

func NewRouter(cache *redis.Client, db *query.Query) (*Server, error) {
	s := &Server{
		db:      db,
		cache:   cache,
		middles: middleware.New(cache, db),
	}

	gin.EnableJsonDecoderUseNumber()
	gin.SetMode(gin.ReleaseMode)
	mux := gin.Default()

	// 限流器
	limiter := rate.NewLimiter(rate.Every(time.Minute), config.Get().Server.Rate)
	mux.Use(func(ctx *gin.Context) {
		if !limiter.Allow() {
			fmt.Println("限流器触发")
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"message": "too many requests",
			})
			return
		}
	})

	valid.Init()

	s.Mux = mux

	s.setWebRouter()

	system := s.Mux.Group("/system")
	{
		system.GET("/health", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"Timestamp": time.Now().UnixMilli(),
				"Host":      ctx.Request.Host,
				"Status":    "ok",
			})
		})
	}

	s.Mux.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": 1, "msg": "网络错误", "errMsg": "错误的路由路径"})
	})
	s.Mux.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": 1, "msg": "网络错误", "errMsg": "错误的请求方式"})
	})

	return s, nil
}

package router

import (
	"hook007/controller/api_web/auth"
	"hook007/controller/api_web/user"

	"github.com/gin-gonic/gin"
)

func (s *Server) setWebRouter() {
	webRouter := s.Mux.Group("/api/v1/web", func(ctx *gin.Context) {
		ctx.Set("_role_", "web")
	})

	needAuth := webRouter.Group("", s.middles.Token)

	authHandler := auth.New(s.cache)
	needAuth.POST("/auth/login", authHandler.Login)
	needAuth.POST("/auth/register", authHandler.Register)

	userHandler := user.New(s.cache)
	needAuth.GET("/user", userHandler.Detail)

}

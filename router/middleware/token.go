package middleware

import (
	"errors"
	"hook007/config"
	"hook007/pkg/response"
	"hook007/pkg/token"

	"github.com/gin-gonic/gin"
)

type LoginUser struct {
	UserId   uint   `json:"userId"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

func (m *middleware) Token(ctx *gin.Context) {
	resp := response.ResData{Ctx: ctx}

	role := ctx.GetString("_role_")
	if role == "" {
		resp.UnauthorizedError(errors.New("非法用户组").Error())
		return
	}

	var (
		cache       = m.cache
		tokenString = ctx.GetHeader("Authorization")
	)

	if tokenString == "" {
		resp.UnauthorizedError(errors.New("token 为空").Error())
		return
	}
	tokenString = tokenString[7:]

	// jwt 验证
	loginUser, err := token.New(config.Get().JWT.Secret).JwtParse(tokenString)
	if err != nil {
		resp.UnauthorizedError(errors.New("登录状态已过期").Error())
		return
	}

	cacheKey := loginUser.Role
	cacheData := cache.Get(ctx, cacheKey).Val()
	if cacheData == "" {
		resp.UnauthorizedError(errors.New("登录状态已过期").Error())
	}

	ctx.Set("_userId_", loginUser.UserId)
}

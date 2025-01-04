package response

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResData struct {
	Code int          `json:"code"`
	Msg  string       `json:"msg"`
	Data any          `json:"data,omitempty"`
	Ctx  *gin.Context `json:"-"`
}

const (
	SUCCESS = 0
	ERROR   = -1
	// ENCRYPT_ERROR = 1000 // 加密错误
	// UnauthorizedError = 1    // 鉴权错误
)

// Success 200
func (r *ResData) Success(data any) {
	r.Code = SUCCESS
	r.Msg = "success"
	r.Data = data
	r.Ctx.PureJSON(http.StatusOK, r)
}
func (r *ResData) Fail(msg string) {
	r.Code = ERROR
	r.Msg = msg
	r.Ctx.JSON(http.StatusOK, r)
}

// BadRequest 400
func (r *ResData) BadRequest(msg string) {
	r.Code = ERROR
	r.Msg = msg
	r.Ctx.AbortWithStatusJSON(http.StatusBadRequest, r)
}

// UnauthorizedError 401
func (r *ResData) UnauthorizedError(msg string) {
	r.Code = ERROR
	r.Msg = msg
	r.Ctx.AbortWithStatusJSON(http.StatusUnauthorized, r)
}

// NotFoundError 404
func (r *ResData) NotFoundError(msg string) {
	r.Code = ERROR
	r.Msg = msg
	r.Ctx.AbortWithStatusJSON(http.StatusNotFound, r)
}

// 500
func (r *ResData) Error(err string) {
	slog.Error("err", "resp err", err)
	r.Code = ERROR
	r.Msg = "内部错误"
	r.Ctx.AbortWithStatusJSON(http.StatusInternalServerError, r)
}

// Result 自定义
func (r *ResData) Result(code int, msg string) {
	r.Code = code
	r.Msg = msg
	r.Ctx.JSON(code, r)
}

package auth

import (
	"encoding/json"
	"fmt"
	"hook007/pkg/response"
	"hook007/pkg/valid"

	"github.com/gin-gonic/gin"
)

func (s *service) Register(ctx *gin.Context) {
	type postSchema struct {
		Name string `json:"name" binding:"required"`
		XXX  uint   `json:"xxx" binding:"omitempty"`
	}

	var (
		body = &postSchema{}
		resp = response.ResData{Ctx: ctx}
	)

	if err := ctx.ShouldBindJSON(body); err != nil {
		bodyJson, _ := json.MarshalIndent(body, "", "\t")
		fmt.Println(string(bodyJson))

		resp.Fail(valid.Translator(err))
		return
	}

	type Result struct{}
	result := &Result{}
	resp.Success(result)
}

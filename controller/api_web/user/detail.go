package user

import (
	"encoding/json"
	"fmt"
	"hook007/pkg/response"
	"hook007/pkg/valid"

	"github.com/gin-gonic/gin"
)

func (s *service) Detail(ctx *gin.Context) {
	type getQuery struct {
		Page    int    `form:"page,default=1" binding:"min=1"`
		PerPage int    `form:"perPage,default=20" binding:"min=1,max=100"`
		Keyword string `form:"keyword"`
	}

	var (
		query = &getQuery{}
		resp  = response.ResData{Ctx: ctx}
	)

	if err := ctx.ShouldBindQuery(query); err != nil {
		queryJson, _ := json.MarshalIndent(query, "", "\t")
		fmt.Println(string(queryJson))

		resp.Fail(valid.Translator(err))
		return
	}

	type Result struct {
		Rows  []interface{} `json:"rows"`
		Count int64         `json:"count"`
	}
	result := &Result{
		// Rows:  list,
		// Count: count,
	}
	resp.Success(result)
}

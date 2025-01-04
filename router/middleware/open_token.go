package middleware

import (
	"hook007/pkg/response"

	"github.com/gin-gonic/gin"
)

func (m *middleware) OpenToken(ctx *gin.Context) {
	var (
		// cache     = m.cache
		resp      = response.ResData{Ctx: ctx}
		accessKey = ctx.GetHeader("AccessKey")
		secret    = ctx.GetHeader("Secret")
	)

	if accessKey == "" || secret == "" {
		resp.UnauthorizedError("access key or secret is empty")
		return
	}

	// q := &model_gen.Qw{}
	// cacheKey := fmt.Sprintf("xcx_mall:login:open:%s", accessKey)
	// cacheData := cache.Get(ctx, cacheKey).Val()
	// if cacheData == "" {
	// 	if err := dal.Qw.
	// 		WithContext(ctx).
	// 		Where(dal.Qw.OpenAccessKey.Eq(accessKey)).
	// 		Scan(q); err != nil {
	// 		resp.UnauthorizedError(err.Error())
	// 		return
	// 	}
	// 	qwByte, _ := json.Marshal(q)
	// 	cache.SetEx(ctx, cacheKey, string(qwByte), time.Minute*10)
	// } else {
	// 	if err := json.Unmarshal([]byte(cacheData), q); err != nil {
	// 		resp.UnauthorizedError("access key or secret is error")
	// 		return
	// 	}
	// }

	// if q.OpenSecret != secret {
	// 	resp.UnauthorizedError("access key or secret is error")
	// 	return
	// }

	// ctx.Set("_open_corpId_", q.CorpID)
}

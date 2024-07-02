package middleware

import (
	"im/pkg/e"
	"im/pkg/util"
	"im/serializer"
	"time"

	"github.com/gin-gonic/gin"
)

// 中间件，实现token鉴权
func AuthCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var claims *util.UserClaims
		var err error
		code := e.Success
		// 获取报文头部的token
		token := ctx.GetHeader("token")
		if token == "" {
			code = e.ErrorAuthToken
		} else {
			claims, err = util.ParseToken(token)
			if err != nil {
				code = e.ErrorAuthToken
			} else if time.Now().Unix() > claims.ExpiresAt {
				// 过期则返回一个token过期错误
				code = e.ErrorTokenTimeOut
			}
		}
		if code != e.Success {
			ctx.JSON(200, serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			})
			ctx.Abort()
			return
		}
		ctx.Set("user_claims", claims)
		ctx.Next()
	}
}

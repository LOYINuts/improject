package router

import (
	"im/api"
	"im/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.StaticFS("/static", http.Dir("./static"))
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
	})
	// 用户操作
	// 用户登录
	r.POST("/user/login", api.UserLogin)
	// 用户注册
	r.POST("/user/register", api.UserRegister)

	auth := r.Group("auth") //需要鉴权
	auth.Use(middleware.AuthCheck())
	{
		// 用户操作
		// 获取用户详细信息
		auth.GET("user/detail", api.UserDetail)
		// 用户更新
		auth.PUT("user/update", api.UserUpdate)
		// 邮件发送
		auth.POST("email/send", api.EmailSend)
		// 邮件验证
		auth.POST("email/varify", api.EmailVarify)
		// 发送接收消息
		auth.GET("websocket/message", api.WebsocketMessage)
		// 获取聊天记录
		auth.GET("chat/list", api.ChatList)
		// 房间操作
	}
	return r
}

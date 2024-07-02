package api

import (
	"im/pkg/util"
	"im/serializer"
	"im/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 用户登录接口
func UserLogin(c *gin.Context) {
	var ul service.UserService
	if err := c.ShouldBind(&ul); err != nil {
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse(err))
	} else {
		res := ul.Login(c.Request.Context())
		c.JSON(http.StatusOK, res)
	}
}

// 用户注册接口
func UserRegister(c *gin.Context) {
	var ur service.UserService
	if err := c.ShouldBind(&ur); err != nil {
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse(err))
	} else {
		res := ur.Register(c.Request.Context())
		c.JSON(http.StatusOK, res)
	}
}

// 用户详情接口
func UserDetail(c *gin.Context) {
	var ud service.UserService
	claims, _ := c.Get("user_claims")
	userClaims := claims.(*util.UserClaims)
	if err := c.ShouldBind(&ud); err != nil {
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse(err))
	} else {
		res := ud.DetailInfo(c.Request.Context(), userClaims.ID)
		c.JSON(http.StatusOK, res)
	}
}

// 用户信息更新接口,只更新NickName,Gender
func UserUpdate(c *gin.Context) {
	claims, _ := c.Get("user_claims")
	userClaims := claims.(*util.UserClaims)
	var uu service.UserService
	if err := c.ShouldBind(&uu); err != nil {
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse(err))
	} else {
		res := uu.Update(c.Request.Context(), userClaims.ID)
		c.JSON(http.StatusOK, res)
	}
}

// 发送邮箱验证码
func EmailSend(c *gin.Context) {
	var es service.EmailService
	claims, _ := c.Get("user_claims")
	userClaims := claims.(*util.UserClaims)
	if err := c.ShouldBind(&es); err != nil {
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse(err))
	} else {
		res := es.Send(c.Request.Context(), userClaims.ID)
		c.JSON(http.StatusOK, res)
	}
}

// 邮箱token验证
func EmailVarify(c *gin.Context) {
	var ev service.EmailService
	if err := c.ShouldBind(&ev); err != nil {
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse(err))
	} else {
		res := ev.Varify(c.Request.Context(), c.GetHeader("token"))
		c.JSON(http.StatusOK, res)
	}
}

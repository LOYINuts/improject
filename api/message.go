package api

import (
	"im/pkg/util"
	"im/serializer"
	"im/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ChatList(c *gin.Context) {
	var cl service.MessageService
	claims := c.MustGet("user_claims")
	userClaims := claims.(*util.UserClaims)
	if err := c.ShouldBind(&cl); err != nil {
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse(err))
	} else {
		res := cl.ChatLog(c.Request.Context(), userClaims.ID)
		c.JSON(http.StatusOK, res)
	}
}

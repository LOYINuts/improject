package api

import (
	"im/pkg/util"
	"im/serializer"
	"im/service"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}
var webconns = make(map[uint]*websocket.Conn)

func WebsocketMessage(c *gin.Context) {
	var wm service.WebsocketService
	// 因为在这里如果Upgrade成功了就会返回200，所以后面的我的c.JSON不起作用
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(serializer.ErrorResponse(err))
		return
	}
	claims := c.MustGet("user_claims")
	userClaims := claims.(*util.UserClaims)
	webconns[userClaims.ID] = conn
	res := wm.Connecting(c.Request.Context(), userClaims.ID, conn, webconns)
	log.Println(res)
	// 连接结束则删除集合里面该id的键值对
	delete(webconns, userClaims.ID)
}

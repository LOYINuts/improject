package service

import (
	"context"
	"im/dao"
	"im/models"
	"im/pkg/e"
	"im/serializer"
	"log"

	"github.com/gorilla/websocket"
)

type WebsocketService struct {
}

func (service *WebsocketService) Connecting(ctx context.Context, uid uint, conn *websocket.Conn, wc map[uint]*websocket.Conn) serializer.Response {
	defer conn.Close()
	code := e.Success
	webDao := dao.NewWebsocketDao(ctx)
	msgDao := dao.NewMessageDao(ctx)
	for {
		msg := new(serializer.MessageVO)
		// 读取消息
		err := conn.ReadJSON(msg)
		if err != nil {
			code = e.ErrorJson
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
		// 判断是否在该房间内,不在则不能发送消息
		_, err = webDao.GetUserRoomByUidRid(uid, msg.RoomId)
		if err != nil {
			log.Printf("UserId:%v doesnt exist in RoomId:%v", uid, msg.RoomId)
			code = e.ErrorUserNotInRoom
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
		// 保存消息
		m := &models.Message{
			UserId:  uid,
			RoomId:  msg.RoomId,
			Content: msg.Message,
		}
		err = msgDao.SaveMessage(m)
		if err != nil {
			code = e.ErrorDataBase
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
		// 获取在特定房间的在线用户
		// 获取在该房间的所有用户
		userRooms, err := webDao.GetUserRoomsByRid(msg.RoomId)
		if err != nil {
			code = e.ErrorDataBase
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
		// 对该房间的在线用户发送信息
		for _, room := range userRooms {
			// 如果该用户在线,则发送该消息
			if cc, ok := wc[room.UserId]; ok {
				err = cc.WriteMessage(websocket.TextMessage, []byte(msg.Message))
				if err != nil {
					code = e.Error
					return serializer.Response{
						Status: code,
						Msg:    e.GetMsg(code),
						Error:  err.Error(),
					}
				}
			}
		}
	}
}

package service

import (
	"context"
	"im/dao"
	"im/models"
	"im/pkg/e"
	"im/serializer"
)

type MessageService struct {
	RoomId uint `json:"room_id" form:"room_id"`
	models.BasePage
}

func (service *MessageService) ChatLog(ctx context.Context, uid uint) serializer.Response {
	code := e.Success
	// roomid为空
	if service.RoomId == 0 {
		code = e.ErrorInvalidParams
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "房间号不能为空!",
		}
	}
	webDao := dao.NewWebsocketDao(ctx)
	// 判断用户是否属于该房间
	_, err := webDao.GetUserRoomByUidRid(uid, service.RoomId)
	if err != nil {
		code = e.ErrorUserNotInRoom
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	// 如果page参数为0则初始化一下
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	if service.PageIndex == 0 {
		service.PageIndex = 1
	}
	msgDao := dao.NewMessageDao(ctx)
	messages, err := msgDao.ListMessages(service.BasePage, service.RoomId)
	if err != nil {
		code = e.ErrorDataBase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildMessageVOs(messages), uint(len(messages)))
}

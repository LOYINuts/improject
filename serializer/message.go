package serializer

import (
	"im/models"
	"time"
)

type MessageVO struct {
	Message  string    `json:"message"`
	RoomId   uint      `json:"room_id"`
	UserId   uint      `json:"user_id"`
	CreateAt time.Time `json:"create_at"`
}

func BuildMessageVO(item *models.Message) *MessageVO {
	return &MessageVO{
		Message:  item.Content,
		RoomId:   item.RoomId,
		UserId:   item.UserId,
		CreateAt: item.CreatedAt,
	}
}

func BuildMessageVOs(items []*models.Message) (messages []*MessageVO) {
	for _, item := range items {
		message := BuildMessageVO(item)
		messages = append(messages, message)
	}
	return
}

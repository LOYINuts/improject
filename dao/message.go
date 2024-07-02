package dao

import (
	"context"
	"im/models"

	"gorm.io/gorm"
)

type MessageDao struct {
	*gorm.DB
}

func NewMessageDao(ctx context.Context) *MessageDao {
	return &MessageDao{NewDBClient(ctx)}
}

func NewMessageDaoByDB(db *gorm.DB) *MessageDao {
	return &MessageDao{db}
}

// 保存信息
func (dao *MessageDao) SaveMessage(msg *models.Message) error {
	return dao.DB.Model(&models.Message{}).Create(&msg).Error
}

// 根据房间id获取聊天记录
func (dao *MessageDao) ListMessages(page models.BasePage, rid uint) (messages []*models.Message, err error) {
	// 按照时间顺序返回聊天记录,离当前时间近的先展示
	err = dao.DB.Model(&models.Message{}).Limit(page.PageSize).Offset((page.PageIndex - 1) * (page.PageSize)).Order("created_at desc").Find(&messages).Error
	return
}

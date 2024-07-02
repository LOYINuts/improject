package dao

import (
	"context"
	"im/models"

	"gorm.io/gorm"
)

type WebsocketDao struct {
	*gorm.DB
}

func NewWebsocketDao(ctx context.Context) *WebsocketDao {
	return &WebsocketDao{NewDBClient(ctx)}
}

func NewWebsocketDaoByDB(db *gorm.DB) *WebsocketDao {
	return &WebsocketDao{db}
}

// 判断该用户是否在该房间
func (dao *WebsocketDao) GetUserRoomByUidRid(uid, rid uint) (UserRoom *models.UserRoom, err error) {
	err = dao.DB.Model(&models.UserRoom{}).Where("user_id=? AND room_id=?", uid, rid).First(&UserRoom).Error
	return
}

// 获取在这个房间的所有用户
func (dao *WebsocketDao) GetUserRoomsByRid(rid uint) (UserRooms []*models.UserRoom, err error) {
	err = dao.DB.Model(&models.UserRoom{}).Where("room_id=?", rid).Find(&UserRooms).Error
	return
}

package models

import "gorm.io/gorm"

type UserRoom struct {
	gorm.Model
	User   User `gorm:"ForeignKey:UserId"`
	UserId uint `gorm:"not null"`
	Room   Room `gorm:"ForeignKey:RoomId"`
	RoomId uint `gorm:"not null"`
}

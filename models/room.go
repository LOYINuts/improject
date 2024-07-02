package models

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	Number string `gorm:"unique"`
	Name   string `gorm:"not null"`
	Info   string
	User   User `gorm:"ForeignKey:UserId"`
	UserId uint `gorm:"not null"`
}

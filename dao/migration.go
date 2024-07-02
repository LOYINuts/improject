package dao

import (
	"fmt"
	"im/models"
)

// 对model进行迁移即建表
func migration() {
	err := _db.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(
		&models.Room{},
		&models.User{},
		&models.Message{},
		&models.UserRoom{},
	)
	if err != nil {
		fmt.Printf("error:%v\n", err)
	}
}

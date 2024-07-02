package dao

import (
	"context"
	"im/models"

	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

// 创建dao
func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDBClient(ctx)}
}

// 复用db创建dao
func NewUserDaoByDB(db *gorm.DB) *UserDao {
	return &UserDao{db}
}

// 判断账户是否存在，只有在不是记录未找到的错误的时候才返回err
func (dao *UserDao) ExistOrNotByAccount(acc string) (user *models.User, exist bool, err error) {
	err = dao.DB.Model(&models.User{}).Where("account=?", acc).First(&user).Error
	// 没有找到，不能用user == nil来判断因为user作为返回值已经进行了初始化
	// 是未找到错误则返回nil错误
	if err == gorm.ErrRecordNotFound {
		return nil, false, nil
	}
	// 如果其他错误则返回
	if err != nil {
		return nil, false, err
	}
	return user, true, nil
}

// 根据Id获取用户
func (dao *UserDao) GetUserById(uid uint) (user *models.User, err error) {
	err = dao.DB.Model(&models.User{}).Where("id=?", uid).First(&user).Error
	return
}

// 创建用户
func (dao *UserDao) CreateUser(user *models.User) error {
	return dao.DB.Model(&models.User{}).Create(&user).Error
}

// 更新用户
func (dao *UserDao) UpdateUser(user *models.User, uid uint) error {
	// 根据struct更新，因为gorm的特性，0值会被忽略，所以我们再单独更新一下可能的0值的字段
	err := dao.DB.Model(&models.User{}).Where("id=?", uid).Updates(&user).Error
	if err != nil {
		return err
	}
	err = dao.DB.Model(&models.User{}).Select("email", "gender").Where("id=?", uid).Updates(&user).Error
	return err
}

// 根据email判断是否有人用过该邮箱
func (dao *UserDao) GetUserCountByEmail(email string) (int64, error) {
	var num int64
	err := dao.DB.Model(&models.User{}).Where("email=?", email).Count(&num).Error
	return num, err
}

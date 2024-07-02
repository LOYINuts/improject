package serializer

import (
	"im/conf"
	"im/models"
)

// 对外界展示的User类
type UserVO struct {
	Account  string `json:"account"`
	NickName string `json:"nick_name"`
	Email    string `json:"email"`
	Gender   uint8  `json:"gender"`
	Avatar   string `json:"avatar"`
}

func BuildUserVO(item *models.User) *UserVO {
	return &UserVO{
		Account:  item.Account,
		NickName: item.NickName,
		Email:    item.Email,
		Gender:   item.Gender,
		Avatar:   conf.Host + conf.HttpPort + conf.AvatarPath + item.Avatar,
	}
}

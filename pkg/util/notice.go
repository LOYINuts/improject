package util

const (
	BindEmail      = 1 //绑定邮箱
	UnbindEmail    = 2 //解绑邮箱
	ChangePassword = 3 //改密
)

var NoticeFlags = map[int]string{
	BindEmail:      "绑定邮箱",
	UnbindEmail:    "解绑邮箱",
	ChangePassword: "改变密码",
}

func GetNotice(code int) string {
	return NoticeFlags[code]
}

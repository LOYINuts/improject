package e

var MsgFlags = map[int]string{
	Success:            "ok",
	Error:              "fail",
	ErrorDataBase:      "数据库操作错误",
	ErrorJson:          "json格式转换错误,或者连接断开",
	ErrorInvalidParams: "参数错误",
	ErrorUserNotFound:  "找不到用户名",
	ErrorPwNotMatch:    "密码错误",
	ErrorAuthToken:     "token错误",
	ErrorTokenTimeOut:  "token过期",
	ErrorEmailExist:    "邮箱已经被绑定!",
	ErrorSendEmail:     "邮件发送失败!",
	ErrorUserNotInRoom: "用户不在该房间",
	ErrorEmailOPType:   "邮箱操作类型错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if !ok {
		return MsgFlags[Error]
	}
	return msg
}

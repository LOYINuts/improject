package e

// 消息code
const (
	Success            = 200  //成功
	Error              = 400  //出错
	ErrorDataBase      = 600  //数据库出错
	ErrorJson          = 9001 //json格式转换错误
	ErrorInvalidParams = 9002 //参数错误
	ErrorUserNotFound  = 9003 //找不到用户
	ErrorPwNotMatch    = 9004 //密码错误
	ErrorAuthToken     = 9005 //token错误
	ErrorTokenTimeOut  = 9006 //token过期
	ErrorEmailExist    = 9007 //email已经被绑定
	ErrorSendEmail     = 9008 //email发送失败
	ErrorUserNotInRoom = 9009 //用户不在该房间
	ErrorEmailOPType   = 9010 //邮箱操作类型错误
)

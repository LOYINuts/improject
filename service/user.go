package service

import (
	"context"
	"im/conf"
	"im/dao"
	"im/models"
	"im/pkg/e"
	"im/pkg/util"
	"im/serializer"
	"time"

	"gopkg.in/gomail.v2"
)

type UserService struct {
	Account  string `json:"account" form:"account"`
	NickName string `json:"nick_name" form:"nick_name"`
	Password string `json:"password" form:"password"`
	Gender   uint8  `json:"gender" form:"gender"`
}

type EmailService struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	// 1:绑定邮箱 2:解绑邮箱 3:改密
	OperationType uint `json:"operation_type" form:"operation_type"`
}

// 登录
func (service *UserService) Login(ctx context.Context) serializer.Response {
	code := e.Success
	if service.Account == "" || service.Password == "" {
		code = e.ErrorInvalidParams
		return serializer.Response{
			Status: code,
			Msg:    "用户名或密码不得为空",
		}
	}
	userDao := dao.NewUserDao(ctx)
	user, exist, err := userDao.ExistOrNotByAccount(service.Account)
	if err != nil {
		code = e.ErrorDataBase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	// 用户不存在
	if !exist {
		code = e.ErrorUserNotFound
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	// 密码匹配
	ok := user.CheckPassword(service.Password)
	if !ok {
		code = e.ErrorPwNotMatch
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	// 登录成功签发token
	token, err := util.GenerateToken(user.ID, user.Account)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data: serializer.TokenData{
			User:  *serializer.BuildUserVO(user),
			Token: token,
		},
	}
}

// 注册
func (service *UserService) Register(ctx context.Context) serializer.Response {
	code := e.Success
	var user models.User
	userDao := dao.NewUserDao(ctx)
	_, exist, err := userDao.ExistOrNotByAccount(service.Account)
	if err != nil {
		code = e.ErrorDataBase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	if exist {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    "用户已存在!",
		}
	}
	user = models.User{
		Account:  service.Account,
		NickName: service.NickName,
		Gender:   0,
		Avatar:   "avatar.png",
	}
	// 设置密码并加密
	err = user.SetPassword(service.Password)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	err = userDao.CreateUser(&user)
	if err != nil {
		code = e.ErrorDataBase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildUserVO(&user),
	}
}

// 展示用户详情
func (service *UserService) DetailInfo(ctx context.Context, uid uint) serializer.Response {
	code := e.Success
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uid)
	if err != nil {
		code = e.ErrorDataBase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildUserVO(user),
	}
}

// 用户更新
func (service *UserService) Update(ctx context.Context, uid uint) serializer.Response {
	code := e.Success
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uid)
	if err != nil {
		code = e.ErrorDataBase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	if service.NickName != "" {
		user.NickName = service.NickName
	}
	if service.Gender != 0 {
		user.Gender = service.Gender
	}
	err = userDao.UpdateUser(user, user.ID)
	if err != nil {
		code = e.ErrorDataBase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildUserVO(user),
	}
}

// 发送验证码
func (service *EmailService) Send(ctx context.Context, uid uint) serializer.Response {
	code := e.Success
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uid)
	if err != nil {
		code = e.ErrorDataBase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	// 发送邮件的目的邮箱
	dstEmail := service.Email
	if service.OperationType == 1 {
		// 如果是绑定邮箱，我们要先看用户目前是否绑定邮箱，如果绑定，则发送给之前的邮箱，否则给service的邮箱
		if user.Email != "" {
			dstEmail = user.Email
		}
		// 顺便判断一下如果没有绑定，判断service的email是否为空
		if dstEmail == "" {
			code = e.ErrorInvalidParams
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Data:   "邮箱不能为空!",
			}
		}
		// 判断是否有用户绑定过该邮箱
		cnt, err := userDao.GetUserCountByEmail(service.Email)
		if err != nil {
			code = e.ErrorDataBase
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
		// 存在绑定该邮箱的账号
		if cnt > 0 {
			code = e.ErrorEmailExist
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
	} else if service.OperationType == 2 {
		//如果是解绑邮箱，则必须用户当前有邮箱
		if user.Email == "" {
			code = e.ErrorInvalidParams
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Data:   "用户未绑定邮箱，无法解绑",
			}
		} else {
			dstEmail = user.Email
		}
	} else if service.OperationType == 3 {
		// 如果是改密，同样用户必须先绑定邮箱才能改密
		if user.Email == "" {
			code = e.ErrorInvalidParams
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Data:   "用户未绑定邮箱，无法改密",
			}
		} else {
			dstEmail = user.Email
		}
	} else {
		// 其他的操作返回错误的操作类型
		code = e.ErrorEmailOPType
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	// 这里生成token还是需要用service的email，因为万一是绑定邮箱则需要用service的邮箱
	token, err := util.GenerateEmailToken(user.ID, service.OperationType, service.Email, service.Password)
	if err != nil {
		code = e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	mailStr := util.GetNotice(int(service.OperationType)) + token
	// 创建新邮件消息
	mail := gomail.NewMessage()
	mail.SetHeader("From", conf.SmtpEmail)
	mail.SetHeader("To", dstEmail)
	mail.SetHeader("Subject", "IMSYSTEM")
	mail.SetBody("text/html", mailStr)
	// 发送
	dial := gomail.NewDialer(conf.SmtpHost, conf.SmtpPort, conf.SmtpEmail, conf.SmtpPass)
	if err = dial.DialAndSend(mail); err != nil {
		code = e.ErrorSendEmail
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

// 验证邮箱操作
func (service *EmailService) Varify(ctx context.Context, token string) serializer.Response {
	code := e.Success
	// 判断token是否为空
	if token == "" {
		code = e.ErrorInvalidParams
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "token不可为空!",
		}
	}
	// 获取claims
	claims, err := util.ParseEmailToken(token)
	if err != nil {
		code = e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	// token过期
	if time.Now().Unix() > claims.ExpiresAt {
		code = e.ErrorTokenTimeOut
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	// 获取claims的数据
	userId := claims.UserID
	email := claims.Email
	password := claims.Password
	opType := claims.OperationType
	userDao := dao.NewUserDao(ctx)
	// 获取用户
	user, err := userDao.GetUserById(userId)
	if err != nil {
		code = e.ErrorDataBase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	if opType == 1 {
		user.Email = email
	} else if opType == 2 {
		user.Email = ""
	} else if opType == 3 {
		if err = user.SetPassword(password); err != nil {
			code = e.Error
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	} else {
		code = e.ErrorEmailOPType
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	// 更新到数据库
	err = userDao.UpdateUser(user, userId)
	if err != nil {
		code = e.ErrorDataBase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   util.GetNotice(int(opType)) + "成功",
	}
}

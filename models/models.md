# Models介绍

## User用户集合

```go
type User struct {
	gorm.Model
	Account  string
	Password string
	NickName string
	Gender   uint8
	Avatar   string
}
```

- gorm.Model包含创建、更新、删除时间以及主键ID
- Account账户
- Password密码
- NickName用户名
- gender性别，0未知，1男，2女
- Avatar头像

## Message消息集合

```go
type Message struct {
	gorm.Model
	UserId  uint
	RoomId  uint
	Content string
}
```

- UserId 用户唯一标识
- RoomId 房间唯一标识
- Content 信息内容



## Room聊天室集合

```go
type Room struct {
	gorm.Model
	Number string
	Name   string
	Info   string
	UserId uint
}
```

- Number 房间号
- Name 房间名
- Info 房间简介
- UserId 房间创建者的用户标识

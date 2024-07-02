package serializer

import "im/pkg/e"

// 返回消息结构
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

type TokenData struct {
	Token string `json:"token"`
	User  UserVO `json:"user"`
}

type DataList struct {
	Item  interface{} `json:"item"`
	Total uint        `json:"total"`
}

func ErrorResponse(err error) *Response {
	return &Response{
		Status: e.Error,
		Msg:    e.GetMsg(e.Error),
		Error:  err.Error(),
	}
}

func BuildListResponse(items interface{}, total uint) Response {
	return Response{
		Status: e.Success,
		Data: DataList{
			Item:  items,
			Total: total,
		},
		Msg: e.GetMsg(e.Success),
	}
}

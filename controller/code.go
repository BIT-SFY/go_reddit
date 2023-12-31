package controller

type ResCode int64

const (
	CodeSuccess = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExit
	CodeInvalidPassword
	CodeServerBusy

	CodeInvalidToken
	CodeNeedLogin

	CodeInvalidCommunityID
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:            "success",
	CodeInvalidParam:       "请求参数错误",
	CodeUserExist:          "用户名已存在",
	CodeUserNotExit:        "用户名不存在",
	CodeInvalidPassword:    "用户名或密码错误",
	CodeServerBusy:         "服务繁忙",
	CodeNeedLogin:          "需要登录",
	CodeInvalidToken:       "无效的token",
	CodeInvalidCommunityID: "群id不存在",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}

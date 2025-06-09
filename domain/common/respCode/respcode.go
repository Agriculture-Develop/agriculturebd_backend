package respCode

type StatusCode int64

// 共有模块
const (
	Success        StatusCode = 200
	Forbidden      StatusCode = 300
	ServerBusy     StatusCode = 400
	RecordNotFound StatusCode = 500
)

// 认证模块
const (
	InvalidParams StatusCode = 1001 + iota
	InvalidParamsFormat
	UserExist
	UserNotExist
	InvalidPassword
	NotLogin
	InvalidCaptcha

	InvalidTokenForm
	InvalidToken
	InvalidTokenExpired

	InvalidDataUpdate
	UserALREADYLocked
)

var Msg = map[StatusCode]string{
	// 共有模块
	Success:        "成功",
	Forbidden:      "权限不足",
	ServerBusy:     "服务繁忙",
	RecordNotFound: "未查询到该记录",

	// 认证模块
	InvalidParams:       "请求参数错误",
	InvalidParamsFormat: "请求参数格式错误",
	UserExist:           "用户已存在",
	UserNotExist:        "用户不存在",
	InvalidPassword:     "手机号或密码错误",
	NotLogin:            "用户未登录",
	InvalidCaptcha:      "手机号或验证码错误",
	InvalidToken:        "无效的Token",
	InvalidTokenForm:    "不合法的token格式",
	InvalidTokenExpired: "Token已过期",
	InvalidDataUpdate:   "不合法的数据更新",
	UserALREADYLocked:   "用户已被锁定",
}

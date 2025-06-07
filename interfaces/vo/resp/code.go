package resp

type StatusCode int64

const (
	// 成功
	CodeSuccess StatusCode = 1000
)

const (
	// 认证模块
	CodeInvalidParams StatusCode = 2001 + iota
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeNotLogin
	CodeInvalidCaptcha

	CodeInvalidTokenForm
	CodeInvalidToken
	CodeInvalidTokenExpired

	CodeInvalidDataUpdate
	CodeUserALREADYLocked
)

// TODO 待规划
const (
	CodeForbidden         StatusCode = 3001
	CodeServerBusy        StatusCode = 4001
	CodeRecordNotFound    StatusCode = 5001
	CodeRateLimitExceeded StatusCode = 6001
)

var Msg = map[StatusCode]string{
	CodeSuccess: "success",

	// 认证模块
	CodeInvalidParams:       "请求参数错误",
	CodeUserExist:           "用户已存在",
	CodeUserNotExist:        "用户不存在",
	CodeInvalidPassword:     "用户名或密码错误",
	CodeNotLogin:            "用户未登录",
	CodeInvalidCaptcha:      "手机号或验证码错误",
	CodeInvalidToken:        "无效的Token",
	CodeInvalidTokenForm:    "不合法的token格式",
	CodeInvalidTokenExpired: "Token已过期",
	CodeInvalidDataUpdate:   "不合法的数据更新",
	CodeUserALREADYLocked:   "用户已被锁定",

	// 其他错误
	CodeForbidden:         "权限不足",
	CodeServerBusy:        "服务繁忙",
	CodeRecordNotFound:    "未查询到该记录",
	CodeRateLimitExceeded: "操作频率过快 ,请稍后再试",
}

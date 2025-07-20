package dto

// 更新用户信息请求
type UpdateUserInfoCtrlDto struct {
	Nickname string `json:"nickname" `
	Role     string `json:"role"`
	Status   string `json:"status"`
}

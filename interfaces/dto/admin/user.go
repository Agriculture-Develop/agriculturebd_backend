package dto

// 更新用户信息请求
type UpdateUserInfoCtrlDto struct {
	Nickname string `json:"nickname" binding:"required"`
	Role     int    `json:"role"  binding:"oneof=0 1 2 3"`
	Status   int    `json:"status"  binding:"oneof=0 1 2"`
}

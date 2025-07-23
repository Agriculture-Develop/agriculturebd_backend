package dto

import "mime/multipart"

// 更新用户信息请求
type UpdateUserInfoCtrlDto struct {
	Nickname string `json:"nickname" `
	Role     string `json:"role"`
	Status   string `json:"status"`
}

type UpdateUserInfoCtrlDtoByUser struct {
	Nickname string                `form:"nickname" `
	Role     string                `form:"role"`
	Avatar   *multipart.FileHeader `form:"avatar"`
}

package vo

// UserVo 用户信息
type UserCtrlVo struct {
	ID         uint   `json:"id"`
	Phone      string `json:"phone"`
	Nickname   string `json:"nickname"`
	AvatarPath string `json:"avatar_path"`
	Role       string `json:"role"`
	Status     string `json:"status"`
	CreatedAt  string `json:"created_at"`
}

// UserListVo 用户列表
type UserListCtrlVo struct {
	Total int64        `json:"total"`
	List  []UserCtrlVo `json:"list"`
}

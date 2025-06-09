package valobj

type UserRole int

const (
	RoleUser       UserRole = iota // 0 - 普通用户
	RoleAdmin                      // 1 - 管理员
	RoleSuperAdmin                 // 2 - 超级管理员
)

func (r UserRole) Desc() string {
	switch r {
	case RoleUser:
		return "用户"
	case RoleAdmin:
		return "管理员"
	case RoleSuperAdmin:
		return "超级管理员"
	default:
		return "未知角色"
	}
}

func (r UserRole) Int() int {
	return int(r)
}

type UserStatus string

const (
	StatusEnabled  UserStatus = "enabled"
	StatusDisabled UserStatus = "disabled"
	StatusBanned   UserStatus = "banned"
)

func (s UserStatus) Desc() string {
	switch s {
	case StatusEnabled:
		return "启用中"
	case StatusDisabled:
		return "已禁用"
	case StatusBanned:
		return "封禁中"
	default:
		return "未知状态"
	}
}

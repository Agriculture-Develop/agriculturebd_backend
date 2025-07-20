package valobj

type UserRole int

const (
	RoleFarmer UserRole = -iota - 1
	RoleSupplier
)

const (
	RoleUser       UserRole = iota // 0 - 普通用户
	RoleAdmin                      // 1 - 管理员
	RoleSuperAdmin                 // 2 - 超级管理员
)

func (r UserRole) Desc() string {
	switch r {
	case RoleUser:
		return "普通用户"
	case RoleAdmin:
		return "管理员"
	case RoleSuperAdmin:
		return "超级管理员"
	case RoleSupplier:
		return "供应商"
	case RoleFarmer:
		return "农户"
	default:
		return "未知角色"
	}
}

func (r UserRole) Int() int {
	return int(r)
}

type UserStatus int

const (
	StatusEnabled  UserStatus = 0
	StatusDisabled UserStatus = 1
	StatusUnknown  UserStatus = 2
)

func (r UserStatus) Int() int {
	return int(r)
}

func (s UserStatus) Desc() string {
	switch s {
	case StatusEnabled:
		return "启用"
	case StatusDisabled:
		return "禁用"
	case StatusUnknown:
		return "未知状态"
	default:
		return "未知状态"
	}
}

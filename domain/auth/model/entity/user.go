package entity

import (
	"regexp"
	"time"

	"github.com/Agriculture-Develop/agriculturebd/domain/auth/model/valobj"
)

type User struct {
	ID       uint
	Phone    string
	Password string
	Role     valobj.UserRole
	Status   valobj.UserStatus

	CreatedAt time.Time
}

func NewUser(phone, password string, role valobj.UserRole) *User {
	return &User{
		Phone:    phone,
		Password: password,
		Role:     role,
		Status:   valobj.StatusEnabled,
	}
}

// IsEnabled 判断是否启用
func (u *User) IsEnabled() bool {
	return u.Status == valobj.StatusEnabled
}

// IsSuperAdmin 判断是否是超级管理员
func (u *User) IsSuperAdmin() bool {
	return u.Role == valobj.RoleSuperAdmin
}

// CheckPassword 密码强度检查
func CheckPassword(password string) bool {
	var (
		hasMinLength = len(password) >= 6
		//hasUpperCase   = regexp.MustCompile(`[A-Z]`).MatchString(password)
		hasLowerCase   = regexp.MustCompile(`[a-z]`).MatchString(password)
		hasNumber      = regexp.MustCompile(`[0-9]`).MatchString(password)
		hasSpecialChar = regexp.MustCompile(`[\~\!\?\@\#\$\%\^\&\*\_\-\+\=\(\)\[\]\{\}\>\<\/\\\|\"\'\.\,\:\;]`).MatchString(password)
	)
	return hasMinLength && hasLowerCase && (hasNumber || hasSpecialChar)
}

// CheckPhone 国内手机号格式校验
func CheckPhone(phone string) bool {
	var phoneRegex = regexp.MustCompile(`^1[3-9]\d{9}$`)
	return phoneRegex.MatchString(phone)
}

package entity

import (
	"github.com/Agriculture-Develop/agriculturebd/domain/user/model/valobj"
	"regexp"
	"time"
)

type User struct {
	ID         uint   `json:"id"`
	Phone      string `json:"phone"`
	Password   string `json:"-"`
	Nickname   string `json:"nickname"`
	AvatarPath string `json:"avatar_path"`
	Role       valobj.UserRole
	Status     valobj.UserStatus
	CreatedAt  time.Time `json:"created_at"`
}

// 创建新用户
func NewUser(phone, password string, role int) *User {
	return &User{
		Phone:    phone,
		Password: password,
		Role:     valobj.RoleUser,
		Status:   valobj.StatusEnabled,
	}
}

// 检查手机号格式
func CheckPhone(phone string) bool {
	var phoneRegex = regexp.MustCompile(`^1[3-9]\d{9}$`)
	return phoneRegex.MatchString(phone)
}

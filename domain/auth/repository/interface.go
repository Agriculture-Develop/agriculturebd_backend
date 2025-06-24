package repository

import (
	"github.com/Agriculture-Develop/agriculturebd/domain/auth/model/entity"
)

type IAuthRepo interface {
	// 根据id获取用户
	GetUserById(id uint) (*entity.User, error)
	// 根据手机号获取用户
	GetUserByPhone(phone string) (*entity.User, error)
	// 创建新用户
	CreateUser(user *entity.User) error
	// 更新用户密码
	UpdateNewPassword(userId uint, newPassword string) error
	// 保存验证码
	SavePhoneCode(phone, code string) error
	// 验证手机验证码
	VerifyPhoneCode(phone, code string) bool
	// 生成并保存token
	GenerateToken(userId uint, role int) (string, error)
}

type ISMSUtils interface {
	SendCaptcha(phone string, code string) error
}

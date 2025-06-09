package repository

import (
	"github.com/Agriculture-Develop/agriculturebd/domain/auth/model/entity"
)

type IAuthRepo interface {
	// 根据手机号获取用户
	GetUserByPhone(phone string) (*entity.User, error)
	// 创建新用户
	CreateUser(user *entity.User) error
	// 保存验证码
	SavePhoneCode(phone, code string) error
	// 验证手机验证码
	VerifyPhoneCode(phone, code string) bool
	// 生成并保存token
	GenerateToken(userId uint) (string, error)
}

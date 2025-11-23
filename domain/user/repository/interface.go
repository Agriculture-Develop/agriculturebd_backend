package repository

import (
	"github.com/Agriculture-Develop/agriculturebd/domain/user/model/entity"
)

type IUserRepo interface {
	// 获取用户列表
	GetUserList(page, count int) ([]*entity.User, error)
	// 获取用户总数
	GetUserCount() (int64, error)
	// 根据ID获取用户
	GetUserById(id uint) (*entity.User, error)
	// 根据角色获取用户ID列表
	GetUserIDsByRole(role int) ([]uint, error)
	// 更新用户信息
	UpdateUser(user *entity.User) error
	// 删除用户
	DeleteUser(id uint) error
}

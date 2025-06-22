package service

import (
	"errors"

	"github.com/Agriculture-Develop/agriculturebd/domain/common/respCode"
	"github.com/Agriculture-Develop/agriculturebd/domain/user/model/valobj"
	"github.com/Agriculture-Develop/agriculturebd/domain/user/repository"
	"github.com/Agriculture-Develop/agriculturebd/domain/user/service/vo"
	"go.uber.org/dig"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type IUserSvc interface {
	// 获取用户列表
	GetUserList(page, count int) (respCode.StatusCode, []vo.UserSvcVo)
	// 更新用户信息
	UpdateUserInfo(userId uint, nickname string, role, status int) respCode.StatusCode
	// 删除用户
	DeleteUser(userId uint) respCode.StatusCode
}

type Svc struct {
	dig.In
	Repo repository.IUserRepo
}

func NewUserSvc(r repository.IUserRepo) IUserSvc {
	return &Svc{Repo: r}
}

func (s *Svc) GetUserList(page, count int) (respCode.StatusCode, []vo.UserSvcVo) {
	// 1. 获取用户列表
	users, err := s.Repo.GetUserList(page, count)
	if err != nil {
		zap.L().Error("GetUserList fail", zap.Error(err))
		return respCode.ServerBusy, nil
	}

	// 2. 转换为VO
	userVos := make([]vo.UserSvcVo, 0, len(users))
	for _, user := range users {
		userVos = append(userVos, vo.UserSvcVo{
			ID:         user.ID,
			Phone:      user.Phone,
			Nickname:   user.Nickname,
			AvatarPath: user.AvatarPath,
			Role:       user.Role.Int(),
			Status:     user.Status.Int(),
			CreatedAt:  user.CreatedAt,
		})
	}

	return respCode.Success, userVos
}

func (s *Svc) UpdateUserInfo(userId uint, nickname string, role, status int) respCode.StatusCode {
	// 1. 获取用户信息
	user, err := s.Repo.GetUserById(userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return respCode.UserNotExist
		}
		zap.L().Error("GetUserById fail", zap.Error(err))
		return respCode.ServerBusy
	}

	// 2. 更新用户信息
	user.Nickname = nickname
	user.Role = valobj.UserRole(role)
	user.Status = valobj.UserStatus(status)
	if err := s.Repo.UpdateUser(user); err != nil {
		zap.L().Error("UpdateUser fail", zap.Error(err))
		return respCode.ServerBusy
	}

	return respCode.Success
}

func (s *Svc) DeleteUser(userId uint) respCode.StatusCode {
	// 1. 检查用户是否存在
	if _, err := s.Repo.GetUserById(userId); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return respCode.UserNotExist
		}
		zap.L().Error("GetUserById fail", zap.Error(err))
		return respCode.ServerBusy
	}

	// 2. 删除用户
	if err := s.Repo.DeleteUser(userId); err != nil {
		zap.L().Error("DeleteUser fail", zap.Error(err))
		return respCode.ServerBusy
	}

	return respCode.Success
}

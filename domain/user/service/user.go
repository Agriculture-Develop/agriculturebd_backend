package service

import (
	"errors"
	"github.com/Agriculture-Develop/agriculturebd/infrastructure/utils/upload"
	"mime/multipart"

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
	UpdateUserInfo(userId uint, nickname string, role, status string) respCode.StatusCode
	// 删除用户
	DeleteUser(userId uint) respCode.StatusCode
	// 获取用户详情
	GetUserDetail(userId uint) (respCode.StatusCode, vo.UserSvcVo)
	// 更新用户头像
	UpdateUserAvatar(userId uint, avatar *multipart.FileHeader) respCode.StatusCode
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
			Role:       user.Role.Desc(),
			Status:     user.Status.Desc(),
			CreatedAt:  user.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return respCode.Success, userVos
}

func (s *Svc) UpdateUserInfo(userId uint, nickname string, role, status string) respCode.StatusCode {
	// 1. 获取用户信息
	user, err := s.Repo.GetUserById(userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return respCode.UserNotExist
		}
		zap.L().Error("GetUserById fail", zap.Error(err))
		return respCode.ServerBusy
	}

	// 校验参数
	if role != "" && valobj.GetUserRole(role) == valobj.RoleUnknown {
		return respCode.InvalidParams
	}

	// 身份校验
	if role != "" && valobj.GetUserRole(role).Int() > 0 && user.Role != valobj.RoleSuperAdmin {
		return respCode.Forbidden
	}

	// 2. 更新用户信息
	user.Nickname = nickname

	if role != "" {
		user.Role = valobj.GetUserRole(role)
	}

	if user.Role == valobj.RoleSuperAdmin && valobj.GetUserStatus(status) != valobj.StatusUnknown {
		user.Status = valobj.GetUserStatus(status)
	}

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

func (s *Svc) GetUserDetail(userId uint) (respCode.StatusCode, vo.UserSvcVo) {
	user, err := s.Repo.GetUserById(userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return respCode.UserNotExist, vo.UserSvcVo{}
		}
		zap.L().Error("GetUserById fail", zap.Error(err))
		return respCode.ServerBusy, vo.UserSvcVo{}
	}
	return respCode.Success, vo.UserSvcVo{
		ID:         user.ID,
		Nickname:   user.Nickname,
		Phone:      user.Phone,
		Role:       user.Role.Desc(),
		AvatarPath: user.AvatarPath,
		Status:     user.Status.Desc(),
		CreatedAt:  user.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}

func (s *Svc) UpdateUserAvatar(userId uint, avatar *multipart.FileHeader) respCode.StatusCode {
	// 查询用户
	user, err := s.Repo.GetUserById(userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return respCode.UserNotExist
		}
		zap.L().Error("[UpdateUserAvatar] GetUserById failed", zap.Error(err))
		return respCode.ServerBusy
	}

	// 删除旧头像
	if user.AvatarPath != "" {
		if err := upload.DeleteFile(user.AvatarPath, "avatar"); err != nil {
			zap.L().Error("[UpdateUserAvatar] DeleteFile failed", zap.Error(err))
			return respCode.ServerBusy
		}
	}

	// 上传新头像
	user.AvatarPath, err = upload.UploadFile(avatar, "avatar")
	if err != nil {
		zap.L().Error("[UpdateUserAvatar] UploadFile failed", zap.Error(err))
		return respCode.ServerBusy
	}

	// 更新用户信息
	if err := s.Repo.UpdateUser(user); err != nil {
		zap.L().Error("[UpdateUserAvatar] UpdateUser failed", zap.Error(err))
		return respCode.ServerBusy
	}

	return respCode.Success
}

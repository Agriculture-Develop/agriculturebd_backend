package user

import (
	"github.com/Agriculture-Develop/agriculturebd/domain/user/model/entity"
	"github.com/Agriculture-Develop/agriculturebd/domain/user/model/valobj"
	"github.com/Agriculture-Develop/agriculturebd/domain/user/repository"
	"github.com/Agriculture-Develop/agriculturebd/infrastructure/dao/model"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

type Repo struct {
	dig.In
	Db *gorm.DB
}

func NewUserRepo(db *gorm.DB) repository.IUserRepo {
	return &Repo{Db: db}
}

func (r *Repo) GetUserList(page, count int) ([]*entity.User, error) {
	var users []model.User
	offset := (page - 1) * count
	if err := r.Db.Offset(offset).Limit(count).Find(&users).Error; err != nil {
		return nil, err
	}

	entityUsers := make([]*entity.User, 0, len(users))
	for _, user := range users {
		entityUsers = append(entityUsers, &entity.User{
			ID:         user.ID,
			Phone:      user.Phone,
			Password:   user.Password,
			Nickname:   user.Nickname,
			AvatarPath: user.AvatarPath,
			Role:       valobj.UserRole(user.Role),
			Status:     valobj.UserStatus(user.Status),
		})
	}
	return entityUsers, nil
}

func (r *Repo) GetUserCount() (int64, error) {
	var count int64
	err := r.Db.Model(&model.User{}).Count(&count).Error
	return count, err
}

func (r *Repo) GetUserById(id uint) (*entity.User, error) {
	var user model.User
	if err := r.Db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &entity.User{
		ID:         user.ID,
		Phone:      user.Phone,
		Password:   user.Password,
		Nickname:   user.Nickname,
		AvatarPath: user.AvatarPath,
		Role:       valobj.UserRole(user.Role),
		Status:     valobj.UserStatus(user.Status),
	}, nil
}

func (r *Repo) UpdateUser(user *entity.User) error {
	return r.Db.Model(&model.User{}).
		Where("id = ?", user.ID).
		Select("Nickname", "Role", "Status").
		Updates(model.User{
			Nickname: user.Nickname,
			Role:     int(user.Role),
			Status:   int(user.Status),
		}).Error
}

func (r *Repo) DeleteUser(id uint) error {
	return r.Db.Delete(&model.User{}, id).Error
}

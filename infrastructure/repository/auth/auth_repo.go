package auth

import (
	"context"
	"errors"
	"github.com/Agriculture-Develop/agriculturebd/domain/auth/model/entity"
	"github.com/Agriculture-Develop/agriculturebd/domain/auth/model/valobj"
	"github.com/Agriculture-Develop/agriculturebd/domain/auth/repository"
	"github.com/Agriculture-Develop/agriculturebd/infrastructure/utils/cache"
	"go.uber.org/dig"
	"time"

	"github.com/Agriculture-Develop/agriculturebd/infrastructure/dao/model"

	"github.com/Agriculture-Develop/agriculturebd/infrastructure/utils/jwt"
	"gorm.io/gorm"
)

type Repo struct {
	dig.In

	Db    *gorm.DB
	Cache *cache.Cache
}

func NewAuthRepo(db *gorm.DB, cache *cache.Cache) repository.IAuthRepo {
	return &Repo{
		Db:    db,
		Cache: cache,
	}
}

func (r *Repo) GetUserById(id uint) (*entity.User, error) {
	var user model.User
	if err := r.Db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &entity.User{
		ID:       user.ID,
		Phone:    user.Phone,
		Password: user.Password,
		Role:     valobj.UserRole(user.Role),
	}, nil

}

func (r *Repo) GetUserByPhone(phone string) (*entity.User, error) {
	var user model.User
	if err := r.Db.Where("phone = ?", phone).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &entity.User{
		ID:       user.ID,
		Phone:    user.Phone,
		Password: user.Password,
		Role:     valobj.UserRole(user.Role),
	}, nil
}

func (r *Repo) CreateUser(user *entity.User) error {
	dbUser := &model.User{
		Phone:    user.Phone,
		Password: user.Password,
		Role:     user.Role.Int(),
		Status:   user.Status.Int(),
	}

	return r.Db.Create(dbUser).Error
}

func (r *Repo) SavePhoneCode(phone, code string) error {
	key := "phone_code:" + phone
	return r.Cache.Set(context.Background(), key, code, 5*time.Minute)
}

func (r *Repo) VerifyPhoneCode(phone, code string) bool {
	key := "phone_code:" + phone
	var savedCode string
	if err := r.Cache.Get(context.Background(), key, &savedCode); err != nil {
		return false
	}
	return savedCode == code
}

func (r *Repo) GenerateToken(userId uint, role int) (string, error) {
	return jwt.GenerateToken(userId, role)
}

func (r *Repo) UpdateNewPassword(userId uint, newPassword string) error {
	return r.Db.Model(&model.User{}).
		Where("id = ?", userId).
		Update("password", newPassword).Error
}

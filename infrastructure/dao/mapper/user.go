package mapper

import (
	"context"
	database "github.com/Agriculture-Develop/agriculturebd/infrastructure/dao/init"
	"github.com/Agriculture-Develop/agriculturebd/infrastructure/dao/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type UserDao struct {
	m *model.User
}

// NewUserDao make user dao
func NewUserDao() *UserDao {
	return &UserDao{m: new(model.User)}
}

func (d *UserDao) db(ctx context.Context) *gorm.DB {
	return database.GetDb().WithContext(ctx)
}

// GetByID get user by id
func (d *UserDao) GetByID(ctx context.Context, id int64) (*model.User, error) {
	var result model.User

	err := d.db(ctx).Model(d.m).Where("id = ?", id).First(&result).Error
	if err != nil {
		return nil, errors.Wrap(err, "user dao get by id error")
	}

	return &result, nil
}

// GetByIDs get user list by ids
func (d *UserDao) GetByIDs(ctx context.Context, ids []int64) ([]*model.User, error) {
	var result []*model.User

	err := d.db(ctx).Model(d.m).Where("id in ?", ids).Find(&result).Error
	if err != nil {
		return nil, errors.Wrap(err, "user dao get by ids error")
	}

	return result, nil
}

// List get user list
func (d *UserDao) List(ctx context.Context, page, size int) ([]*model.User, error) {
	var (
		list   []*model.User
		total  int64
		offset = (page - 1) * size
		limit  = size
	)

	err := d.db(ctx).Model(d.m).Count(&total).Offset(offset).Limit(limit).Find(&list).Error
	if err != nil {
		return nil, errors.Wrap(err, "user dao get page error")
	}

	//pageTotal := int32(math.Ceil(float64(total) / float64(page.PageSize)))
	//pageResp := &base.PaginationResp{Page: page.Page, PageSize: page.PageSize, Total: int32(total), PageTotal: pageTotal}

	return list, nil
}

// Create insert a user record
func (d *UserDao) Create(ctx context.Context, m *model.User) error {
	err := d.db(ctx).Model(d.m).Create(m).Error
	if err != nil {
		return errors.Wrap(err, "user dao create error")
	}

	return nil
}

// Update user by id
func (d *UserDao) Update(ctx context.Context, id int64, m *model.User) error {
	err := d.db(ctx).Model(d.m).Where("id = ?", id).Updates(m).Error
	if err != nil {
		return errors.Wrap(err, "user dao update error")
	}

	return nil
}

// Delete delete user by id
func (d *UserDao) Delete(ctx context.Context, id int64) error {
	err := d.db(ctx).Model(d.m).Where("id = ?", id).Delete(d.m).Error
	if err != nil {
		return errors.Wrap(err, "user dao delete error")
	}

	return nil
}

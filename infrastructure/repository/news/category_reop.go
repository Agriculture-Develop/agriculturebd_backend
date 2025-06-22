package news

import (
	"github.com/Agriculture-Develop/agriculturebd/domain/news/entity"
	"github.com/Agriculture-Develop/agriculturebd/domain/news/repository"
	"github.com/Agriculture-Develop/agriculturebd/infrastructure/dao/model"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

type CategoryRepo struct {
	dig.In
	Db *gorm.DB
}

func NewNewsCategoryRepo(db *gorm.DB) repository.INewsCategoryRepo {
	return &CategoryRepo{Db: db}
}

func (r *CategoryRepo) Create(category *entity.NewsCategory) error {
	dbCategory := &model.NewsCategories{
		Name:        category.Name,
		Description: category.Description,
		SortOrder:   category.SortOrder,
	}
	return r.Db.Create(dbCategory).Error
}

func (r *CategoryRepo) Update(category *entity.NewsCategory) error {
	return r.Db.Model(&model.NewsCategories{}).Where("id = ?", category.ID).Updates(map[string]interface{}{
		"name":        category.Name,
		"description": category.Description,
		"sort_order":  category.SortOrder,
	}).Error
}

func (r *CategoryRepo) Delete(id uint) error {
	return r.Db.Delete(&model.NewsCategories{}, id).Error
}

func (r *CategoryRepo) GetByID(id uint) (*entity.NewsCategory, error) {
	var dbCategory model.NewsCategories
	if err := r.Db.First(&dbCategory, id).Error; err != nil {
		return nil, err
	}
	return &entity.NewsCategory{
		ID:          dbCategory.ID,
		Name:        dbCategory.Name,
		Description: dbCategory.Description,
		SortOrder:   dbCategory.SortOrder,
		CreatedAt:   dbCategory.CreatedAt,
		UpdatedAt:   dbCategory.UpdatedAt,
	}, nil
}

func (r *CategoryRepo) List() ([]*entity.NewsCategory, error) {
	var dbCategories []model.NewsCategories
	if err := r.Db.Order("sort_order asc, created_at desc").Find(&dbCategories).Error; err != nil {
		return nil, err
	}

	categories := make([]*entity.NewsCategory, 0, len(dbCategories))
	for _, c := range dbCategories {
		categories = append(categories, &entity.NewsCategory{
			ID:          c.ID,
			Name:        c.Name,
			Description: c.Description,
			SortOrder:   c.SortOrder,
			CreatedAt:   c.CreatedAt,
			UpdatedAt:   c.UpdatedAt,
		})
	}

	return categories, nil
}

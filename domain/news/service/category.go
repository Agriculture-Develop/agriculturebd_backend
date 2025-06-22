package service

import (
	"errors"

	"github.com/Agriculture-Develop/agriculturebd/domain/common/respCode"
	"github.com/Agriculture-Develop/agriculturebd/domain/news/entity"
	"github.com/Agriculture-Develop/agriculturebd/domain/news/repository"
	"github.com/Agriculture-Develop/agriculturebd/domain/news/service/dto"
	"github.com/Agriculture-Develop/agriculturebd/domain/news/service/vo"
	"go.uber.org/dig"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type INewsCategorySvc interface {
	CreateCategory(dto dto.CategoryCreateSvcDTO) respCode.StatusCode
	UpdateCategory(id uint, dto dto.CategoryUpdateSvcDTO) respCode.StatusCode
	DeleteCategory(id uint) respCode.StatusCode
	ListCategories() (respCode.StatusCode, *vo.CategoryListSvcVO)
}

type NewsCategorySvc struct {
	dig.In
	CategoryRepo repository.INewsCategoryRepo
}

func NewNewsCategoryService(categoryRepo repository.INewsCategoryRepo) INewsCategorySvc {
	return &NewsCategorySvc{
		CategoryRepo: categoryRepo,
	}
}

// CreateCategory 创建分类
func (s *NewsCategorySvc) CreateCategory(dto dto.CategoryCreateSvcDTO) respCode.StatusCode {
	if dto.Name == "" {
		return respCode.InvalidParamsFormat
	}
	category := &entity.NewsCategory{
		Name:        dto.Name,
		Description: dto.Description,
		SortOrder:   dto.SortOrder,
	}
	if err := s.CategoryRepo.Create(category); err != nil {
		zap.L().Error("CreateCategory fail", zap.Error(err))
		return respCode.ServerBusy
	}
	return respCode.Success
}

// UpdateCategory 更新分类
func (s *NewsCategorySvc) UpdateCategory(id uint, dto dto.CategoryUpdateSvcDTO) respCode.StatusCode {
	category, err := s.CategoryRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return respCode.RecordNotFound
		}
		zap.L().Error("GetCategoryById fail", zap.Error(err))
		return respCode.ServerBusy
	}
	if dto.Name != "" {
		category.Name = dto.Name
	}
	if dto.Description != "" {
		category.Description = dto.Description
	}
	category.SortOrder = dto.SortOrder
	if err := s.CategoryRepo.Update(category); err != nil {
		zap.L().Error("UpdateCategory fail", zap.Error(err))
		return respCode.ServerBusy
	}
	return respCode.Success
}

// DeleteCategory 删除分类
func (s *NewsCategorySvc) DeleteCategory(id uint) respCode.StatusCode {
	if _, err := s.CategoryRepo.GetByID(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return respCode.RecordNotFound
		}
		zap.L().Error("GetCategoryById for delete fail", zap.Error(err))
		return respCode.ServerBusy
	}

	if err := s.CategoryRepo.Delete(id); err != nil {
		zap.L().Error("DeleteCategory fail", zap.Error(err))
		return respCode.ServerBusy
	}
	return respCode.Success
}

// ListCategories 获取分类列表
func (s *NewsCategorySvc) ListCategories() (respCode.StatusCode, *vo.CategoryListSvcVO) {
	categories, err := s.CategoryRepo.List()
	if err != nil {
		zap.L().Error("ListCategories fail", zap.Error(err))
		return respCode.ServerBusy, nil
	}
	voList := make([]vo.CategorySvcVO, 0, len(categories))
	for _, c := range categories {
		voList = append(voList, vo.CategorySvcVO{
			ID:          c.ID,
			Name:        c.Name,
			Description: c.Description,
			SortOrder:   c.SortOrder,
			CreatedAt:   c.CreatedAt,
			UpdatedAt:   c.UpdatedAt,
		})
	}
	return respCode.Success, &vo.CategoryListSvcVO{List: voList}
}

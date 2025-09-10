package service

import (
	"errors"
	"github.com/Agriculture-Develop/agriculturebd/domain/common/respCode"
	"github.com/Agriculture-Develop/agriculturebd/domain/news/entity"
	"github.com/Agriculture-Develop/agriculturebd/domain/news/repository"
	"github.com/Agriculture-Develop/agriculturebd/domain/news/service/dto"
	"github.com/Agriculture-Develop/agriculturebd/domain/news/service/vo"
	"github.com/Agriculture-Develop/agriculturebd/infrastructure/utils/upload"
	"go.uber.org/dig"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type INewsSvc interface {
	CreateNews(dto dto.NewsCreateSvcDTO) respCode.StatusCode
	UpdateNews(id uint, dto dto.NewsUpdateSvcDTO) respCode.StatusCode
	UpdateNewsStatus(id uint, status string) respCode.StatusCode
	GetNewsDetail(id uint) (respCode.StatusCode, *vo.NewsDetailSvcVO)
	ListNews(filter dto.NewsListFilterSvcDTO) (respCode.StatusCode, *vo.NewsListSvcVO)
	DeleteNews(id uint) respCode.StatusCode
}

type NewsSvc struct {
	dig.In
	NewsRepo     repository.INewsRepo
	CategoryRepo repository.INewsCategoryRepo
}

func NewNewsService(newsRepo repository.INewsRepo, categoryRepo repository.INewsCategoryRepo) INewsSvc {
	return &NewsSvc{
		NewsRepo:     newsRepo,
		CategoryRepo: categoryRepo,
	}
}

// CreateNews 创建新闻
func (s *NewsSvc) CreateNews(dto dto.NewsCreateSvcDTO) respCode.StatusCode {
	// 0. 校验状态参数
	if dto.Status != string(entity.StatusDraft) && dto.Status != string(entity.StatusReviewing) {
		return respCode.InvalidParams
	}

	// 1. 检查分类是否存在
	if _, err := s.CategoryRepo.GetByID(dto.CategoryID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return respCode.RecordNotFound
		}
		zap.L().Error("GetCategoryById fail", zap.Error(err))
		return respCode.ServerBusy
	}

	// 4. 创建新闻实体
	news := &entity.News{
		Title:      dto.Title,
		CategoryID: dto.CategoryID,
		Abstract:   dto.Abstract,
		Keyword:    dto.Keyword,
		Source:     dto.Source,
		Content:    dto.Content,
		Type:       entity.NewsType(dto.Types),
		Status:     entity.NewsStatus(dto.Status),
		FilesURL:   dto.FilesURL,
		CoverURL:   dto.CoverURL,
		UserID:     dto.UserID,
	}

	// 5. 保存到数据库
	if err := s.NewsRepo.Create(news); err != nil {
		zap.L().Error("CreateNews fail", zap.Error(err))
		return respCode.ServerBusy
	}

	return respCode.Success
}

// UpdateNews 更新新闻
func (s *NewsSvc) UpdateNews(id uint, dto dto.NewsUpdateSvcDTO) respCode.StatusCode {
	// 1. 获取新闻信息
	news, err := s.NewsRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return respCode.RecordNotFound
		}
		zap.L().Error("GetNewsById fail", zap.Error(err))
		return respCode.ServerBusy
	}

	// 2. 检查分类是否存在（即使没有变也检查）
	if _, err := s.CategoryRepo.GetByID(dto.CategoryID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return respCode.RecordNotFound
		}
		zap.L().Error("GetCategoryById fail", zap.Error(err))
		return respCode.ServerBusy
	}

	// 3. 删除旧封面文件
	if news.CoverURL != "" {
		if err := upload.DeleteFile(news.CoverURL, "news"); err != nil {
			zap.L().Warn("Delete old cover image fail", zap.Error(err))
		}
	}

	// 5. 全量字段更新
	news.Title = dto.Title
	news.CategoryID = dto.CategoryID
	news.Abstract = dto.Abstract
	news.Type = entity.NewsType(dto.Type)
	news.Source = dto.Source
	news.Content = dto.Content
	news.CoverURL = dto.CoverURL
	news.Status = entity.NewsStatus(dto.Status)
	news.UserID = dto.UserID
	news.FilesURL = dto.FilesURL
	news.Keyword = dto.Keyword

	// 8. 保存更新
	if err := s.NewsRepo.Update(news); err != nil {
		zap.L().Error("UpdateNews fail", zap.Error(err))
		return respCode.ServerBusy
	}

	return respCode.Success
}

// UpdateNewsStatus 更新新闻状态
func (s *NewsSvc) UpdateNewsStatus(id uint, status string) respCode.StatusCode {
	// 1. 检查新闻是否存在
	if _, err := s.NewsRepo.GetByID(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return respCode.RecordNotFound
		}
		zap.L().Error("GetNewsById fail", zap.Error(err))
		return respCode.ServerBusy
	}

	// 2. 验证状态值
	validStatus := map[string]bool{
		"未提交":   true,
		"审核中":   true,
		"审核已通过": true,
		"审核已驳回": true,
		"未发布":   true,
		"已发布":   true,
		"已下线":   true,
	}
	if !validStatus[status] {
		return respCode.InvalidParamsFormat
	}

	// 3. 更新状态
	if err := s.NewsRepo.UpdateStatus(id, status); err != nil {
		zap.L().Error("UpdateNewsStatus fail", zap.Error(err))
		return respCode.ServerBusy
	}

	return respCode.Success
}

// GetNewsDetail 获取新闻详情
func (s *NewsSvc) GetNewsDetail(id uint) (respCode.StatusCode, *vo.NewsDetailSvcVO) {
	// 1. 获取新闻信息
	news, err := s.NewsRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return respCode.RecordNotFound, nil
		}
		zap.L().Error("GetNewsById fail", zap.Error(err))
		return respCode.ServerBusy, nil
	}

	// 2. 获取分类信息
	category, err := s.CategoryRepo.GetByID(news.CategoryID)
	if err != nil {
		zap.L().Error("GetCategoryById fail", zap.Error(err))
		return respCode.ServerBusy, nil
	}

	// 获取用户作者信息
	author, err := s.NewsRepo.GetAuthorByID(news.UserID)

	// 5. 构建返回VO
	newsVO := &vo.NewsDetailSvcVO{
		ID:         news.ID,
		Title:      news.Title,
		CategoryID: news.CategoryID,
		Category:   category.Name,
		Abstract:   news.Abstract,
		Keyword:    news.Keyword,
		Source:     news.Source,
		Content:    news.Content,
		CoverURL:   news.CoverURL,
		FilesURL:   news.FilesURL,
		Status:     string(news.Status),
		Author:     author,
		Types:      string(news.Type),
		CreatedAt:  news.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:  news.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	return respCode.Success, newsVO
}

// ListNews 获取新闻列表
func (s *NewsSvc) ListNews(filter dto.NewsListFilterSvcDTO) (respCode.StatusCode, *vo.NewsListSvcVO) {
	// 1. 设置默认分页参数
	if filter.Page <= 0 {
		filter.Page = 1
	}
	if filter.PageSize <= 0 {
		filter.PageSize = 10
	}

	// 2. 构建筛选条件
	repoFilter := repository.NewsListFilter{
		Title:    filter.Title,
		Author:   filter.Author,
		Status:   filter.Status,
		Page:     filter.Page,
		PageSize: filter.PageSize,
	}

	// 3. 获取新闻列表
	newsList, total, err := s.NewsRepo.List(repoFilter)
	if err != nil {
		zap.L().Error("ListNews fail", zap.Error(err))
		return respCode.ServerBusy, nil
	}

	// 4. 获取所有分类信息（用于批量查询）
	categories, err := s.CategoryRepo.List()
	if err != nil {
		zap.L().Error("ListCategories fail", zap.Error(err))
		return respCode.ServerBusy, nil
	}

	// 5. 构建分类映射
	categoryMap := make(map[uint]string)
	for _, cat := range categories {
		categoryMap[cat.ID] = cat.Name
	}

	// 6. 转换为VO
	newsVOs := make([]vo.NewsDetailSvcVO, 0, len(newsList))
	for _, news := range newsList {

		// 获取用户作者信息
		var author string
		author, err = s.NewsRepo.GetAuthorByID(news.UserID)

		newsVO := vo.NewsDetailSvcVO{
			ID:         news.ID,
			Title:      news.Title,
			CategoryID: news.CategoryID,
			Category:   categoryMap[news.CategoryID],
			Abstract:   news.Abstract,
			Keyword:    news.Keyword,
			Source:     news.Source,
			Content:    news.Content,
			CoverURL:   news.CoverURL,
			FilesURL:   news.FilesURL,
			Status:     string(news.Status),
			Author:     author,
			Types:      string(news.Type),
			CreatedAt:  news.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:  news.UpdatedAt.Format("2006-01-02 15:04:05"),
		}
		newsVOs = append(newsVOs, newsVO)
	}

	// 7. 构建返回结果
	result := &vo.NewsListSvcVO{
		Total: int(total),
		List:  newsVOs,
	}

	return respCode.Success, result
}

// 删除新闻
func (s *NewsSvc) DeleteNews(id uint) respCode.StatusCode {
	// 删除本地文件
	news, err := s.NewsRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return respCode.RecordNotFound
		}
		zap.L().Error("GetNewsById fail", zap.Error(err))
		return respCode.ServerBusy
	}

	if news.CoverURL != "" {
		err := upload.DeleteFile(news.CoverURL, "news")
		if err != nil {
			return respCode.ServerBusy
		}
	}

	for _, fileURL := range news.FilesURL {
		err := upload.DeleteFile(fileURL, "news")
		if err != nil {
			return respCode.ServerBusy
		}
	}

	err = s.NewsRepo.Delete(id)
	if err != nil {
		zap.L().Error("DeleteNews fail", zap.Error(err))
		return respCode.ServerBusy
	}

	return respCode.Success
}

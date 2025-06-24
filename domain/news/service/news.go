package service

import (
	"encoding/json"
	"errors"
	"github.com/Agriculture-Develop/agriculturebd/domain/common/respCode"
	"github.com/Agriculture-Develop/agriculturebd/domain/news/entity"
	"github.com/Agriculture-Develop/agriculturebd/domain/news/repository"
	"github.com/Agriculture-Develop/agriculturebd/domain/news/service/dto"
	"github.com/Agriculture-Develop/agriculturebd/domain/news/service/vo"
	"go.uber.org/dig"
	"go.uber.org/zap"
	"gorm.io/datatypes"
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
	// 1. 参数校验
	if dto.Title == "" || dto.Content == "" {
		return respCode.InvalidParamsFormat
	}

	// 2. 检查分类是否存在
	if _, err := s.CategoryRepo.GetByID(dto.CategoryID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return respCode.RecordNotFound
		}
		zap.L().Error("GetCategoryById fail", zap.Error(err))
		return respCode.ServerBusy
	}

	// 3. 转换关键词为JSON
	keywordJSON, err := json.Marshal(dto.Keyword)
	if err != nil {
		zap.L().Error("Marshal keyword fail", zap.Error(err))
		return respCode.ServerBusy
	}

	// 4. 转换文件URL为JSON
	filesURLJSON, err := json.Marshal(dto.FilesURL)
	if err != nil {
		zap.L().Error("Marshal files_url fail", zap.Error(err))
		return respCode.ServerBusy
	}

	// 5. 创建新闻实体
	news := &entity.News{
		Title:      dto.Title,
		CategoryID: dto.CategoryID,
		Abstract:   dto.Abstract,
		Keyword:    datatypes.JSON(keywordJSON),
		Source:     dto.Source,
		Content:    dto.Content,
		Status:     entity.NewsStatus(dto.Status),
		FilesURL:   datatypes.JSON(filesURLJSON),
		CoverURL:   dto.CoverURL,
		UserID:     dto.UserID,
	}

	// 6. 保存到数据库
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

	// 2. 检查分类是否存在（如果更新了分类）
	if dto.CategoryID != 0 && dto.CategoryID != news.CategoryID {
		if _, err := s.CategoryRepo.GetByID(dto.CategoryID); err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return respCode.RecordNotFound
			}
			zap.L().Error("GetCategoryById fail", zap.Error(err))
			return respCode.ServerBusy
		}
	}

	// 3. 更新字段
	if dto.Title != "" {
		news.Title = dto.Title
	}
	if dto.CategoryID != 0 {
		news.CategoryID = dto.CategoryID
	}
	if dto.Abstract != "" {
		news.Abstract = dto.Abstract
	}
	if dto.Source != "" {
		news.Source = dto.Source
	}
	if dto.Content != "" {
		news.Content = dto.Content
	}
	if dto.CoverURL != "" {
		news.CoverURL = dto.CoverURL
	}

	// 4. 更新关键词
	if dto.Keyword != nil {
		keywordJSON, err := json.Marshal(dto.Keyword)
		if err != nil {
			zap.L().Error("Marshal keyword fail", zap.Error(err))
			return respCode.ServerBusy
		}
		news.Keyword = datatypes.JSON(keywordJSON)
	}

	// 5. 更新文件URL
	if dto.FilesURL != nil {
		filesURLJSON, err := json.Marshal(dto.FilesURL)
		if err != nil {
			zap.L().Error("Marshal files_url fail", zap.Error(err))
			return respCode.ServerBusy
		}
		news.FilesURL = datatypes.JSON(filesURLJSON)
	}

	// 6. 保存更新
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
		"draft":       true,
		"reviewing":   true,
		"approved":    true,
		"rejected":    true,
		"unpublished": true,
		"published":   true,
		"offline":     true,
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

	// 3. 解析关键词
	var keywords []string
	if news.Keyword != nil {
		if err := json.Unmarshal(news.Keyword, &keywords); err != nil {
			zap.L().Error("Unmarshal keyword fail", zap.Error(err))
			return respCode.ServerBusy, nil
		}
	}

	// 获取用户作者信息
	author, err := s.NewsRepo.GetAuthorByID(news.UserID)

	// 4. 解析文件URL
	var filesURL []string
	if news.FilesURL != nil {
		if err := json.Unmarshal(news.FilesURL, &filesURL); err != nil {
			zap.L().Error("Unmarshal files_url fail", zap.Error(err))
			return respCode.ServerBusy, nil
		}
	}

	// 5. 构建返回VO
	newsVO := &vo.NewsDetailSvcVO{
		ID:         news.ID,
		Title:      news.Title,
		CategoryID: news.CategoryID,
		Category:   category.Name,
		Abstract:   news.Abstract,
		Keyword:    keywords,
		Source:     news.Source,
		Content:    news.Content,
		CoverURL:   news.CoverURL,
		FilesURL:   filesURL,
		Status:     string(news.Status),
		Author:     author,
		CreatedAt:  news.CreatedAt,
		UpdatedAt:  news.UpdatedAt,
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
		// 解析关键词
		var keywords []string
		if news.Keyword != nil {
			if err := json.Unmarshal(news.Keyword, &keywords); err != nil {
				zap.L().Error("Unmarshal keyword fail", zap.Error(err))
				continue
			}
		}

		// 获取用户作者信息
		var author string
		author, err = s.NewsRepo.GetAuthorByID(news.UserID)

		// 解析文件URL
		var filesURL []string
		if news.FilesURL != nil {
			if err := json.Unmarshal(news.FilesURL, &filesURL); err != nil {
				zap.L().Error("Unmarshal files_url fail", zap.Error(err))
				continue
			}
		}

		newsVO := vo.NewsDetailSvcVO{
			ID:         news.ID,
			Title:      news.Title,
			CategoryID: news.CategoryID,
			Category:   categoryMap[news.CategoryID],
			Abstract:   news.Abstract,
			Keyword:    keywords,
			Source:     news.Source,
			Content:    news.Content,
			CoverURL:   news.CoverURL,
			FilesURL:   filesURL,
			Status:     string(news.Status),
			Author:     author,
			CreatedAt:  news.CreatedAt,
			UpdatedAt:  news.UpdatedAt,
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
	err := s.NewsRepo.Delete(id)
	if err != nil {
		zap.L().Error("DeleteNews fail", zap.Error(err))
		return respCode.ServerBusy
	}

	return respCode.Success
}

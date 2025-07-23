package news

import (
	"encoding/json"
	"github.com/Agriculture-Develop/agriculturebd/domain/news/entity"
	"github.com/Agriculture-Develop/agriculturebd/domain/news/repository"
	"github.com/Agriculture-Develop/agriculturebd/infrastructure/dao/model"
	"go.uber.org/dig"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type NewsRepo struct {
	dig.In
	Db *gorm.DB
}

func NewNewsRepo(db *gorm.DB) repository.INewsRepo {
	return &NewsRepo{Db: db}
}

func (r *NewsRepo) Create(news *entity.News) error {
	keywordJSON, _ := json.Marshal(news.Keyword)
	filesURLJSON, _ := json.Marshal(news.FilesURL)

	dbNews := &model.News{
		Title:      news.Title,
		Abstract:   news.Abstract,
		Keyword:    datatypes.JSON(keywordJSON),
		Source:     news.Source,
		Content:    news.Content,
		Status:     string(news.Status),
		Comment:    news.Comment,
		FilesURL:   datatypes.JSON(filesURLJSON),
		CoverURL:   news.CoverURL,
		UserID:     news.UserID,
		CategoryID: news.CategoryID,
	}
	return r.Db.Create(dbNews).Error
}

func (r *NewsRepo) Update(news *entity.News) error {
	keywordJSON, _ := json.Marshal(news.Keyword)
	filesURLJSON, _ := json.Marshal(news.FilesURL)

	return r.Db.Model(&model.News{}).Where("id = ?", news.ID).Updates(map[string]interface{}{
		"title":       news.Title,
		"abstract":    news.Abstract,
		"keyword":     datatypes.JSON(keywordJSON),
		"source":      news.Source,
		"content":     news.Content,
		"comment":     news.Comment,
		"files_url":   datatypes.JSON(filesURLJSON),
		"cover_url":   news.CoverURL,
		"category_id": news.CategoryID,
	}).Error
}

func (r *NewsRepo) UpdateStatus(id uint, status string) error {
	return r.Db.Model(&model.News{}).Where("id = ?", id).Update("status", status).Error
}

func (r *NewsRepo) GetByID(id uint) (*entity.News, error) {
	var dbNews model.News
	if err := r.Db.First(&dbNews, id).Error; err != nil {
		return nil, err
	}

	return &entity.News{
		ID:          dbNews.ID,
		Title:       dbNews.Title,
		Abstract:    dbNews.Abstract,
		Keyword:     dbNews.Keyword,
		Source:      dbNews.Source,
		Content:     dbNews.Content,
		Status:      entity.NewsStatus(dbNews.Status),
		Comment:     dbNews.Comment,
		FilesURL:    dbNews.FilesURL,
		CoverURL:    dbNews.CoverURL,
		UserID:      dbNews.UserID,
		CategoryID:  dbNews.CategoryID,
		Type:        entity.NewsType(dbNews.Type),
		PublishedAt: dbNews.PublishedAt,
		CreatedAt:   dbNews.CreatedAt,
		UpdatedAt:   dbNews.UpdatedAt,
	}, nil
}
func (r *NewsRepo) List(filter repository.NewsListFilter) ([]*entity.News, int64, error) {
	var dbNews []model.News
	var total int64

	query := r.Db.Model(&model.News{})

	if filter.Title != "" {
		query = query.Where("news.title LIKE ?", "%"+filter.Title+"%")
	}
	if filter.Status != "" {
		query = query.Where("news.status = ?", filter.Status)
	}
	if filter.Author != "" {
		query = query.Joins("JOIN users ON users.id = news.user_id").
			Where("users.nickname LIKE ?", "%"+filter.Author+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (filter.Page - 1) * filter.PageSize
	if err := query.Order("news.created_at DESC").Offset(offset).Limit(filter.PageSize).Find(&dbNews).Error; err != nil {
		return nil, 0, err
	}

	newsList := make([]*entity.News, 0, len(dbNews))
	for _, n := range dbNews {
		newsList = append(newsList, &entity.News{
			ID:          n.ID,
			Title:       n.Title,
			Abstract:    n.Abstract,
			Keyword:     n.Keyword,
			Source:      n.Source,
			Status:      entity.NewsStatus(n.Status),
			Content:     n.Content,
			UserID:      n.UserID,
			Type:        entity.NewsType(n.Type),
			CategoryID:  n.CategoryID,
			PublishedAt: n.PublishedAt,
			CreatedAt:   n.CreatedAt,
		})
	}

	return newsList, total, nil
}

func (r *NewsRepo) GetAuthorByID(id uint) (name string, err error) {
	dbUser := new(model.User)
	err = r.Db.Where("id = ?", id).First(dbUser).Error
	return dbUser.Nickname, err
}

// 删除新闻
func (r *NewsRepo) Delete(id uint) error {
	return r.Db.Delete(&model.News{}, id).Error
}

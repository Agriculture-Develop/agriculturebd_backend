package repository

import "github.com/Agriculture-Develop/agriculturebd/domain/news/entity"

type INewsRepo interface {
	Create(news *entity.News) error
	Update(news *entity.News) error
	UpdateStatus(id uint, status string) error
	GetByID(id uint) (*entity.News, error)
	List(filter NewsListFilter) ([]*entity.News, int64, error)
}

type NewsListFilter struct {
	Title    string
	Author   string
	Status   string
	Page     int
	PageSize int
}

type INewsCategoryRepo interface {
	Create(category *entity.NewsCategory) error
	Update(category *entity.NewsCategory) error
	Delete(id uint) error
	GetByID(id uint) (*entity.NewsCategory, error)
	List() ([]*entity.NewsCategory, error)
}
